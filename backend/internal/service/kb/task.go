package kb

import (
	"context"
	"errors"
	"flowing/global"
	"flowing/internal/dag"
	"flowing/internal/docprocess"
	"flowing/internal/model/ai"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/repository"
	"flowing/internal/repository/vector"
	"flowing/internal/util"
	"fmt"
	"log/slog"
	"slices"
	"sync"
	"time"

	"github.com/cloudwego/eino/schema"
)

const defaultSliceSize int64 = 1024

var (
	runningTasks = sync.Map{}
)

type ParseTask struct {
	kb.Task
	knowledgeBase *kb.KnowledgeBase
	datasource    *monitor.Datasource
	doc           *kb.Document
	kgModelDetail *ai.ProviderModelDetail
	extractKg     bool
}

func CancelTask(ctx context.Context, id int64) error {
	return repository.Tx(ctx, func(c context.Context) error {
		var task *kb.Task
		// 检查任务是否存在
		err := repository.DB(ctx).Model(&kb.Task{}).
			Where("id = ?", id).
			Where("history = ?", false).
			First(&task).
			Error
		if err != nil {
			return global.NewError(500, "任务不存在", err)
		}
		// 删除任务
		if err := repository.DB(ctx).Delete(&kb.Task{}, "id = ?", id).Error; err != nil {
			return global.NewError(500, "取消任务失败", err)
		}
		// 取消任务
		if run, ok := runningTasks.Load(id); ok {
			run.(*dag.GraphRun).Cancel()
		}
		return nil
	})
}

func CreateTask(ctx context.Context, req kb.CreateTaskReq) error {
	if req.SliceSize <= 0 {
		req.SliceSize = defaultSliceSize
	}
	task := kb.Task{
		DocumentId:         req.DocumentId,   // 文档ID
		Status:             kb.TaskStatusNew, // 新任务
		SliceSize:          req.SliceSize,    // 切片大小
		EndTime:            time.Now(),
		SlicingStartTime:   time.Now(),
		SlicingEndTime:     time.Now(),
		EmbeddingStartTime: time.Now(),
		EmbeddingEndTime:   time.Now(),
	}
	return repository.Tx(ctx, func(c context.Context) error {
		// 获取文档
		doc, _ := kb.GetDocument(c, req.DocumentId)
		if doc == nil {
			return global.NewError(400, "文档不存在", nil)
		}
		task.KnowledgeBaseId = doc.KnowledgeBaseId
		// 判断是否有未完成的任务
		var oldTask int64
		_ = repository.DB(c).Model(&task).
			Where("document_id = ?", req.DocumentId).
			Where("status NOT IN (?)", []kb.TaskStatus{kb.TaskStatusFailed, kb.TaskStatusSuccess}).
			Where("history = ?", false).
			Count(&oldTask).
			Error
		if oldTask > 0 {
			return global.NewError(400, "文档处理任务已存在", nil)
		}
		// 将旧任务设置为历史任务
		err := repository.DB(c).Model(&kb.Task{}).
			Where("document_id = ?", req.DocumentId).
			Where("history = ?", false).
			Update("history", true).
			Error
		if err != nil {
			return global.NewError(500, "创建任务失败", fmt.Errorf("设置历史任务失败: %w", err))
		}

		// 创建任务
		if err := kb.CreateTask(c, &task); err != nil {
			return global.NewError(500, "创建任务失败", err)
		}

		pt := &ParseTask{
			Task: task,
		}
		// 获取知识库
		base, _ := kb.GetKnowledgeBase(c, task.KnowledgeBaseId)
		if base == nil {
			return global.NewError(400, "知识库不存在", nil)
		}
		pt.knowledgeBase = base
		// 获取数据源
		datasource, _ := monitor.GetDatasource(c, base.DatasourceId)
		if datasource == nil {
			return global.NewError(400, "数据源不存在", nil)
		}
		pt.datasource = datasource
		pt.doc = doc
		// 获取知识抽取模型
		pt.extractKg = *pt.knowledgeBase.EnableKnowledgeGraph
		if *pt.knowledgeBase.EnableKnowledgeGraph {
			kgModelDetail, err := ai.GetProviderModelDetail(ctx, pt.knowledgeBase.ChatModel)
			if err != nil {
				return global.NewError(400, "知识抽取模型不存在", err)
			}
			pt.kgModelDetail = kgModelDetail
		}
		// 运行任务
		if err := runTask(pt); err != nil {
			return global.NewError(500, "创建任务失败", err)
		}
		return nil
	})
}

func runTask(p *ParseTask) error {
	// 三个节点：切片、嵌入、存储
	chain := dag.NewChain().AddNode(
		dag.NewNode("slicing", "slicing", p.slicingNodeFunc),
		dag.NewNode("embedding", "embedding", p.embeddingNodeFunc),
		dag.NewNode("store", "store", p.storeNodeFunc),
	)
	// 知识图谱节点
	if p.extractKg {
		chain.AddNode(dag.NewNode("knowledge_graph", "knowledge_graph", p.extractKnowledgeGraph))
	}
	if err := chain.Compile(); err != nil {
		return err
	}
	run := dag.NewChainRun(chain)
	if err := run.Run(
		dag.WithNonBlocking(),
		dag.WithParallelNum(1),
		dag.WithPanicHandler(p.panicHandler),
		dag.WithCallback(p.callback),
		dag.WithWorkerPool(global.WorkerPool()),
	); err != nil {
		return err
	}
	runningTasks.Store(p.Task.Id, run)
	return nil
}

// panicHandler 任务运行异常处理
func (p *ParseTask) panicHandler(err error) {
	p.Status = kb.TaskStatusFailed
	p.ErrorMessage = fmt.Sprintf("任务运行失败: %v", err)
	p.EndTime = time.Now()
	e := kb.UpdateTask(context.Background(), &p.Task)
	if e != nil {
		slog.Error("更新任务失败", "err", e, "task", p.Task)
	}
}

// callback dag任务回调，用于更新数据库任务状态
func (p *ParseTask) callback(event dag.CallbackEvent) {
	if event.Type == dag.EventTypeNodeEnd {
		nodeInfo := event.Data.(dag.ExecInfo)
		switch nodeInfo.Id {
		case "slicing":
			p.Task.Status = kb.TaskStatusEmbedding
			p.Task.SlicingEndTime = time.Now()
			p.Task.EmbeddingStartTime = time.Now()
			p.Task.SliceCount = int64(event.GlobalData["chunk_count"].(int))
		case "store":
			p.Task.Status = kb.TaskStatusSuccess
			p.Task.EmbeddingEndTime = time.Now()
			p.Task.EndTime = time.Now()
		}
		if err := kb.UpdateTask(context.Background(), &p.Task); err != nil {
			slog.Error("更新任务失败", "err", err, "task", p.Task)
		}
	}
}

// slicingNodeFunc 文档切片节点
func (p *ParseTask) slicingNodeFunc(ctx context.Context, _ dag.Node) (result dag.NodeFuncReturn) {
	if !slices.Contains(docprocess.SupportedTypes, p.doc.Type) {
		panic(errors.New("不支持的文件类型"))
	}
	// 切片开始，更新任务状态
	p.Status = kb.TaskStatusSlicing
	p.SlicingStartTime = time.Now()
	if err := kb.UpdateTask(ctx, &p.Task); err != nil {
		panic(fmt.Errorf("更新任务状态失败: %w", err))
	}
	// 下载文档
	content, err := repository.File().Download(ctx, KnowledgeBaseCollectionName(p.doc.KnowledgeBaseId), p.doc.Uri)
	if err != nil {
		panic(fmt.Errorf("下载文件失败: %w", err))
	}
	// 读取文档内容
	docs, err := docprocess.ParseDocument(ctx, p.doc, content)
	if err != nil {
		panic(fmt.Errorf("读取文档失败: %w", err))
	}
	// 文档切片
	// TODO 切片设置
	chunks, err := docprocess.SplitDocument(ctx, docs, p.doc, &docprocess.SplitOption{
		ChunkSize:   int(p.SliceSize),
		OverlapSize: 128,
		Separators:  []string{"\n\n", "\n"},
	})
	if err != nil {
		panic(fmt.Errorf("切片文档失败: %w", err))
	}
	result.Output = map[string]any{
		"sliced_chunks": chunks,
		"chunk_count":   len(chunks),
	}
	return
}

// embeddingNodeFunc 文档嵌入节点
func (p *ParseTask) embeddingNodeFunc(ctx context.Context, _ dag.Node) (result dag.NodeFuncReturn) {
	chunks := ctx.Value("sliced_chunks").([]*schema.Document)
	// 获取知识库嵌入模型
	modelDetail, err := GetEmbeddingModel(ctx, p.knowledgeBase.EmbeddingModel)
	if err != nil {
		panic(fmt.Errorf("获取嵌入模型失败: %w", err))
	}
	// 嵌入
	embeddings, err := docprocess.Embed(ctx, chunks, &docprocess.EmbedOption{
		ProviderType: modelDetail.ProviderType,
		ModelName:    modelDetail.ModelName,
		Config:       modelDetail.ProviderConfig,
	})
	if err != nil {
		panic(fmt.Errorf("嵌入文档失败: %w", err))
	}
	result.Output = map[string]any{
		"embeddings": embeddings,
	}
	return
}

// storeNodeFunc 向量存储节点
func (p *ParseTask) storeNodeFunc(ctx context.Context, _ dag.Node) (result dag.NodeFuncReturn) {
	embeddings := ctx.Value("embeddings").([][]float64)
	chunks := ctx.Value("sliced_chunks").([]*schema.Document)
	vectors := make([]vector.Slice, 0, len(embeddings))
	for i, embedding := range embeddings {
		// 转换为float32
		// TODO 是否存在精度损失？
		embeddingF32 := make([]float32, len(embedding))
		for i, v := range embedding {
			embeddingF32[i] = float32(v)
		}
		vectors = append(vectors, &kb.Slice{
			SliceId:         chunks[i].ID,
			DocumentId:      p.DocumentId,
			KnowledgeBaseId: p.KnowledgeBaseId,
			SliceContent:    chunks[i].Content,
			Embedding:       embeddingF32,
			Index:           int64(i),
		})
	}

	// 向量存储
	vectorStore, err := repository.NewVectorStore(p.datasource)
	if err != nil {
		panic(fmt.Errorf("获取向量存储失败: %w", err))
	}
	err = vectorStore.Add(ctx, KnowledgeBaseCollectionName(p.doc.KnowledgeBaseId), vectors)
	if err != nil {
		panic(fmt.Errorf("存储向量失败: %w", err))
	}
	return
}

func (p *ParseTask) extractKnowledgeGraph(ctx context.Context, _ dag.Node) (result dag.NodeFuncReturn) {
	// 知识抽取
	chatModel, err := util.GetChatModel(ctx, *p.kgModelDetail, false, true)
	if err != nil {
		panic(fmt.Errorf("获取聊天模型失败: %w", err))
	}
	builder, err := docprocess.NewKnowledgeGraphBuilder(repository.Config().KnowledgeGraph.EntityPrompt,
		repository.Config().KnowledgeGraph.RelationPrompt, chatModel)
	if err != nil {
		panic(fmt.Errorf("创建知识图谱构建器失败: %w", err))
	}
	chunks := ctx.Value("sliced_chunks").([]*schema.Document)
	// 知识抽取
	graph, err := builder.Build(ctx, chunks)
	if err != nil {
		panic(fmt.Errorf("构建知识图谱失败: %w", err))
	}
	// TODO 知识图谱存储
	result.Output = map[string]any{
		"knowledge_graph": graph,
	}
	return
}

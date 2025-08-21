package kb

import (
	"context"
	"errors"
	"flowing/global"
	"flowing/internal/docprocess"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/repository"
	"flowing/internal/repository/vector"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"time"

	"github.com/cloudwego/eino/schema"
)

const defaultSliceSize int64 = 1024

type ParseTask struct {
	kb.Task
	knowledgeBase *kb.KnowledgeBase
	datasource    *monitor.Datasource
	doc           *kb.Document
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
		// 调度器取消任务
		pt := ParseTask{Task: *task}
		global.Worker().Cancel(pt.ID())
		// 删除任务
		if err := repository.DB(ctx).Delete(&kb.Task{}, "id = ?", id).Error; err != nil {
			return global.NewError(500, "取消任务失败", err)
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
		if err := global.Worker().Submit(pt); err != nil {
			return global.NewError(500, "创建任务失败", err)
		}
		return nil
	})
}

func (p *ParseTask) Run(ctx context.Context) error {
	defer func() {
		// 处理任务过程的panic
		if err := recover(); err != nil {
			p.Status = kb.TaskStatusFailed
			p.ErrorMessage = fmt.Sprintf("任务运行失败: %v", err)
			p.EndTime = time.Now()
			e := kb.UpdateTask(ctx, &p.Task)
			if e != nil {
				slog.Error("更新任务失败", "err", e, "task", p.Task)
			}
		}
	}()
	if !slices.Contains(docprocess.SupportedTypes, p.doc.Type) {
		panic(errors.New("不支持的文件类型"))
	}
	chunks := p.slicing(ctx)
	// 切片返回空，任务已经取消
	if chunks == nil {
		return nil
	}
	embeddings := p.embedding(ctx, chunks)
	// 嵌入返回空，任务已经取消
	if embeddings == nil {
		return nil
	}

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

	// 存储
	stores := p.store(ctx, vectors)
	// 存储返回空，任务已经取消
	if stores == nil {
		return nil
	}
	return nil
}

func (p *ParseTask) slicing(ctx context.Context) []*schema.Document {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("切片错误: %w", err.(error)))
		}
	}()
	// 切片开始，更新任务状态
	p.Status = kb.TaskStatusSlicing
	p.SlicingStartTime = time.Now()
	if err := kb.UpdateTask(ctx, &p.Task); err != nil {
		panic(fmt.Errorf("更新任务状态失败: %w", err))
	}
	// 开始切片前，检查任务是否已取消
	if p.cancelled(ctx) {
		return nil
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
	chunks, err := docprocess.SplitDocument(ctx, docs, p.doc, &docprocess.SplitOption{
		ChunkSize:   int(p.SliceSize),
		OverlapSize: 128,
		Separators:  []string{"\n\n", "\n"},
	})
	if err != nil {
		panic(fmt.Errorf("切片文档失败: %w", err))
	}
	// 切片结束，更新任务状态
	p.SliceCount = int64(len(chunks))
	p.SlicingEndTime = time.Now()
	if err := kb.UpdateTask(ctx, &p.Task); err != nil {
		panic(fmt.Errorf("更新任务状态失败: %w", err))
	}
	return chunks
}

func (p *ParseTask) embedding(ctx context.Context, chunks []*schema.Document) [][]float64 {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("嵌入错误: %w", err.(error)))
		}
	}()
	// 切片开始，更新任务状态
	p.Status = kb.TaskStatusEmbedding
	p.SlicingStartTime = time.Now()
	if err := kb.UpdateTask(ctx, &p.Task); err != nil {
		panic(fmt.Errorf("更新任务状态失败: %w", err))
	}
	// 开始切片前，检查任务是否已取消
	if p.cancelled(ctx) {
		return nil
	}
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
	// 嵌入结束，更新任务状态
	p.EmbeddingEndTime = time.Now()
	if err := kb.UpdateTask(ctx, &p.Task); err != nil {
		panic(fmt.Errorf("更新任务状态失败: %w", err))
	}
	return embeddings
}

func (p *ParseTask) store(ctx context.Context, slices []vector.Slice) []vector.Slice {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("存储错误: %w", err.(error)))
		}
	}()
	// 存储前，检查任务是否已取消
	if p.cancelled(ctx) {
		return nil
	}
	// 向量存储
	vectorStore, err := repository.NewVectorStore(p.datasource)
	if err != nil {
		panic(fmt.Errorf("获取向量存储失败: %w", err))
	}
	err = vectorStore.Add(ctx, KnowledgeBaseCollectionName(p.doc.KnowledgeBaseId), slices)
	if err != nil {
		panic(fmt.Errorf("存储向量失败: %w", err))
	}
	// 存储结束，更新任务状态
	p.Status = kb.TaskStatusSuccess
	p.EndTime = time.Now()
	if err := kb.UpdateTask(ctx, &p.Task); err != nil {
		panic(fmt.Errorf("更新任务状态失败: %w", err))
	}
	return slices
}

func mockWait(ctx context.Context) {
	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	select {
	case <-timer.C:
	}
}

func (p *ParseTask) cancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func (p *ParseTask) ID() string {
	return strconv.FormatInt(p.Id, 10)
}

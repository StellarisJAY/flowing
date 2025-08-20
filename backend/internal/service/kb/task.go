package kb

import (
	"context"
	"flowing/global"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/repository"
	"flowing/internal/repository/vector"
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

const defaultSliceSize int64 = 512

type ParseTask struct {
	kb.Task
	knowledgeBase *kb.KnowledgeBase
	datasource    *monitor.Datasource
	doc           *kb.Document
}

func CancelTask(ctx context.Context, id int64) error {
	var task *kb.Task
	err := repository.DB(ctx).Model(&kb.Task{}).
		Where("id = ?", id).
		Where("history = ?", false).
		First(&task).
		Error
	if err != nil {
		return global.NewError(500, "任务不存在", err)
	}
	pt := ParseTask{Task: *task}
	global.Worker().Cancel(pt.ID())
	task.Status = kb.TaskStatusCancelled
	task.EndTime = time.Now()
	if err := repository.DB(ctx).Save(&task).Error; err != nil {
		return global.NewError(500, "取消任务失败", err)
	}
	return nil
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
			Where("status NOT IN (?)", []kb.TaskStatus{kb.TaskStatusCancelled, kb.TaskStatusFailed}).
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
		if err := recover(); err != nil {
			p.Status = kb.TaskStatusFailed
			p.ErrorMessage = fmt.Sprintf("任务运行失败: %v", err)
			p.EndTime = time.Now()
			e := kb.UpdateNotCancelledTask(ctx, &p.Task)
			if e != nil {
				slog.Error("更新任务失败", "err", e, "task", p.Task)
			}
		}
	}()
	// 1. 切片
	p.SlicingStartTime = time.Now()
	if p.cancelled(ctx) {
		return nil
	}
	slices, err := p.slicing(ctx)
	if p.cancelled(ctx) {
		return nil
	}
	p.Status = kb.TaskStatusEmbedding
	p.SlicingEndTime = time.Now()
	p.SliceCount = int64(len(slices))
	if err != nil {
		p.Status = kb.TaskStatusFailed
		p.ErrorMessage = err.Error()
		p.EndTime = time.Now()
	}
	// 1.1 更新任务
	if err := kb.UpdateNotCancelledTask(ctx, &p.Task); err != nil {
		return fmt.Errorf("更新任务失败: %w", err)
	}
	if p.Status == kb.TaskStatusFailed {
		return nil
	}
	// 2. 嵌入
	p.EmbeddingStartTime = time.Now()
	if p.cancelled(ctx) {
		return nil
	}
	err = p.embedding(ctx, slices)
	if p.cancelled(ctx) {
		return nil
	}
	p.EmbeddingEndTime = time.Now()
	if err != nil {
		p.Status = kb.TaskStatusFailed
		p.ErrorMessage = err.Error()
		p.EndTime = time.Now()
	}
	// 2.1 更新任务
	if err := kb.UpdateNotCancelledTask(ctx, &p.Task); err != nil {
		return fmt.Errorf("更新任务失败: %w", err)
	}
	if p.Status == kb.TaskStatusFailed {
		return nil
	}
	// 3. 写入向量库
	if p.cancelled(ctx) {
		return nil
	}
	err = p.store(ctx, slices)
	if p.cancelled(ctx) {
		return nil
	}
	p.EndTime = time.Now()
	if err != nil {
		p.Status = kb.TaskStatusFailed
		p.ErrorMessage = err.Error()
	}
	if err := kb.UpdateNotCancelledTask(ctx, &p.Task); err != nil {
		return fmt.Errorf("更新任务失败: %w", err)
	}
	return nil
}

func (p *ParseTask) slicing(ctx context.Context) ([]vector.Slice, error) {
	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	<-timer.C
	return nil, nil
}

func (p *ParseTask) embedding(ctx context.Context, slices []vector.Slice) error {
	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	<-timer.C
	return nil
}

func (p *ParseTask) cancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func (p *ParseTask) store(ctx context.Context, slices []vector.Slice) error {
	panic("not implement")
}

func (p *ParseTask) ID() string {
	return strconv.FormatInt(p.Id, 10)
}

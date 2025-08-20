package kb

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"time"
)

type TaskStatus string

const (
	TaskStatusNew       TaskStatus = "new"       // 新创建
	TaskStatusSlicing   TaskStatus = "slicing"   // 切片中
	TaskStatusEmbedding TaskStatus = "embedding" // 嵌入中
	TaskStatusSuccess   TaskStatus = "success"   // 成功
	TaskStatusFailed    TaskStatus = "failed"    // 失败
	TaskStatusCancelled TaskStatus = "cancelled" // 已取消
)

// Task 知识库文档解析任务
type Task struct {
	common.BaseModel
	DocumentId         int64      `json:"documentId" gorm:"column:document_id;type:bigint;not null;"`
	KnowledgeBaseId    int64      `json:"knowledgeBaseId" gorm:"column:knowledge_base_id;type:bigint;not null;"` // 知识库ID
	Status             TaskStatus `json:"status" gorm:"column:status;type:varchar(16);not null;"`                // 任务状态
	SlicingStartTime   time.Time  `json:"slicingStartTime" gorm:"column:slicing_start_time;"`                    // 切片开始时间
	SlicingEndTime     time.Time  `json:"slicingEndTime" gorm:"column:slicing_end_time;"`                        // 切片结束时间
	EmbeddingStartTime time.Time  `json:"embeddingStartTime" gorm:"column:embedding_start_time;"`                // 嵌入开始时间
	EmbeddingEndTime   time.Time  `json:"embeddingEndTime" gorm:"column:embedding_end_time;"`                    // 嵌入结束时间
	EndTime            time.Time  `json:"endTime" gorm:"column:end_time;"`                                       // 结束时间
	ErrorMessage       string     `json:"errorMessage" gorm:"column:error_message;type:text;"`                   // 错误信息
	SliceCount         int64      `json:"sliceCount" gorm:"column:slice_count;type:bigint;not null;"`            // 切片数量
	SliceSize          int64      `json:"sliceSize" gorm:"column:slice_size;type:bigint;not null;"`              // 切片大小
	History            *bool      `json:"history" gorm:"column:history;type:bool;not null;default:false;"`       // 历史任务
}

func (t *Task) TableName() string {
	return "kb_task"
}

type CreateTaskReq struct {
	DocumentId int64 `json:"documentId,string" binding:"required"` // 文档ID
	SliceSize  int64 `json:"sliceSize,string"`                     // 解析时指定切片大小，默认512
	// TODO 更多解析设置
}

func CreateTask(ctx context.Context, task *Task) error {
	return repository.DB(ctx).Create(task).Error
}

func UpdateTask(ctx context.Context, task *Task) error {
	return repository.DB(ctx).Model(task).Updates(task).Error
}

func ListTasks(ctx context.Context, docIds []int64) ([]*Task, error) {
	var list []*Task
	if err := repository.DB(ctx).Model(&Task{}).
		Where("document_id IN ?", docIds).
		Where("history = ?", false).
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func UpdateNotCancelledTask(ctx context.Context, task *Task) error {
	return repository.DB(ctx).Model(task).
		Where("id = ?", task.Id).
		Where("status != ?", TaskStatusCancelled).
		Where("status != ?", TaskStatusSuccess).
		Updates(task).Error
}

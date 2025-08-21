package kb

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
	"mime/multipart"
)

type Document struct {
	common.BaseModel
	OriginalName    string `json:"originalName" gorm:"type:varchar(255);not null"`     // 文档名称
	Type            string `json:"type" gorm:"type:varchar(255);not null"`             // 文档后缀名类型
	MIMEType        string `json:"mimeType" gorm:"type:varchar(255);not null"`         // 文档MIME类型
	KnowledgeBaseId int64  `json:"knowledgeBaseId,string" gorm:"type:bigint;not null"` // 知识库ID
	Uri             string `json:"uri" gorm:"type:varchar(255);not null"`              // 文档URI （如果是MINIO存储，则是文档的对象key）
	Size            int64  `json:"size" gorm:"type:int;not null"`                      // 文档大小

	Task *Task `json:"task" gorm:"-"`
}

func (d *Document) TableName() string {
	return "ai_knowledge_document"
}

type DocumentQuery struct {
	common.BaseQuery
	KnowledgeBaseId int64  `json:"knowledgeBaseId,string" form:"knowledgeBaseId" binding:"required"`
	Name            string `json:"name" form:"name"`
}

type UploadDocumentReq struct {
	KnowledgeBaseId int64          `json:"knowledgeBaseId,string"`
	FileName        string         `json:"fileName"`
	Size            int64          `json:"size"`
	FileObj         multipart.File `json:"-"`
	ContentType     string         `json:"contentType"`
}

type RenameDocumentReq struct {
	Id           int64  `json:"id,string" binding:"required"`
	OriginalName string `json:"originalName" binding:"required"`
}

func ListDocument(ctx context.Context, query DocumentQuery) ([]*Document, int64, error) {
	var list []*Document
	var total int64
	d := repository.DB(ctx).Model(&Document{}).
		Where("knowledge_base_id = ?", query.KnowledgeBaseId)
	if query.Name != "" {
		d = d.Where("original_name LIKE ?", "%"+query.Name+"%")
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func GetDocument(ctx context.Context, id int64) (*Document, error) {
	var document *Document
	if err := repository.DB(ctx).First(&document, id).Error; err != nil {
		return nil, err
	}
	return document, nil
}

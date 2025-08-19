package kb

import (
	"context"
	"flowing/global"
	"flowing/internal/model/kb"
	"flowing/internal/repository"
)

func ListDocument(ctx context.Context, query kb.DocumentQuery) ([]*kb.Document, int64, error) {
	res, total, err := kb.ListDocument(ctx, query)
	if err != nil {
		return nil, 0, global.NewError(500, "获取文档列表失败", err)
	}
	return res, total, nil
}

func UploadDocument(ctx context.Context, req kb.UploadDocumentReq) error {
	doc := &kb.Document{
		OriginalName:    req.FileName,
		KnowledgeBaseId: req.KnowledgeBaseId,
		Size:            req.Size,
		Type:            req.ContentType,
	}
	return repository.Tx(ctx, func(c context.Context) error {
		knowledgeBase, _ := kb.GetKnowledgeBase(ctx, req.KnowledgeBaseId)
		if knowledgeBase == nil {
			return global.NewError(500, "知识库不存在", nil)
		}
		uri, err := repository.File().Upload(c, req.FileObj)
		if err != nil {
			return global.NewError(500, "上传文件失败", err)
		}
		doc.Uri = uri
		if err := repository.DB(c).Create(doc).Error; err != nil {
			return global.NewError(500, "创建文档失败", err)
		}
		return nil
	})
}

func RenameDocument(ctx context.Context, req kb.RenameDocumentReq) error {
	err := repository.DB(ctx).Model(&kb.Document{}).Where("id = ?", req.Id).UpdateColumns(map[string]any{
		"original_name": req.OriginalName,
	}).Error
	if err != nil {
		return global.NewError(500, "重命名文档失败", err)
	}
	return nil
}

func DeleteDocument(ctx context.Context, id int64) error {
	return repository.Tx(ctx, func(c context.Context) error {
		var doc *kb.Document
		if err := repository.DB(c).First(&doc, id).Error; err != nil {
			return global.NewError(500, "删除文档失败", err)
		}
		if err := repository.DB(c).Delete(&kb.Document{}, "id = ?", id).Error; err != nil {
			return global.NewError(500, "删除文档失败", err)
		}
		if err := repository.File().Delete(c, doc.Uri); err != nil {
			return global.NewError(500, "删除文档失败", err)
		}
		return nil
	})
}

func GetDownloadURL(ctx context.Context, id int64) (string, error) {
	var doc kb.Document
	if err := repository.DB(ctx).First(&doc, "id = ?", id).Error; err != nil {
		return "", global.NewError(500, "获取文档失败", err)
	}
	url, err := repository.File().TempDownloadURL(ctx, doc.Uri)
	if err != nil {
		return "", global.NewError(500, "生成下载URL失败", err)
	}
	return url, nil
}

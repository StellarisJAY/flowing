package kb

import (
	"context"
	"flowing/global"
	"flowing/internal/model/common"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/model/system"
	"flowing/internal/repository"
	"flowing/internal/repository/vector"
	"strings"
)

func ListDocument(ctx context.Context, query kb.DocumentQuery) ([]*kb.Document, int64, error) {
	res, total, err := kb.ListDocument(ctx, query)
	if err != nil {
		return nil, 0, global.NewError(500, "获取文档列表失败", err)
	}
	ids := make([]int64, len(res))
	for i, item := range res {
		ids[i] = item.Id
	}
	tasks, err := kb.ListTasks(ctx, ids)
	if err != nil {
		return nil, 0, global.NewError(500, "获取文档列表失败", err)
	}
	taskMap := make(map[int64]*kb.Task)
	for _, item := range tasks {
		taskMap[item.DocumentId] = item
	}
	for _, item := range res {
		if task, ok := taskMap[item.Id]; ok {
			item.Task = task
		}
	}
	return res, total, nil
}

func UploadDocument(ctx context.Context, req kb.UploadDocumentReq) error {
	index := strings.LastIndex(req.FileName, ".")
	if index == -1 {
		return global.NewError(500, "文件名称格式错误，必须包含文件扩展名", nil)
	}
	doc := &kb.Document{
		BaseModel: common.BaseModel{
			CreateBy: ctx.Value(global.ContextKeyUser).(system.User).Id,
		},
		OriginalName:    req.FileName,
		KnowledgeBaseId: req.KnowledgeBaseId,
		Size:            req.Size,
		MIMEType:        req.ContentType,
		Type:            req.FileName[index+1:],
	}
	return repository.Tx(ctx, func(c context.Context) error {
		knowledgeBase, _ := kb.GetKnowledgeBase(ctx, req.KnowledgeBaseId)
		if knowledgeBase == nil {
			return global.NewError(500, "知识库不存在", nil)
		}
		uri, err := repository.File().Upload(c, KnowledgeBaseCollectionName(knowledgeBase.Id), req.FileObj)
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
		if err := repository.File().Delete(c, KnowledgeBaseCollectionName(doc.KnowledgeBaseId), doc.Uri); err != nil {
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
	url, err := repository.File().TempDownloadURL(ctx, KnowledgeBaseCollectionName(doc.KnowledgeBaseId), doc.Uri)
	if err != nil {
		return "", global.NewError(500, "生成下载URL失败", err)
	}
	return url, nil
}

func ListChunks(ctx context.Context, query vector.ListSliceQuery) ([]vector.QueriedSlice, int64, error) {
	var doc *kb.Document
	if err := repository.DB(ctx).First(&doc, "id = ?", query.DocId).Error; err != nil {
		return nil, 0, global.NewError(500, "获取文档失败", err)
	}
	knowledgeBase, err := kb.GetKnowledgeBase(ctx, doc.KnowledgeBaseId)
	if err != nil {
		return nil, 0, global.NewError(500, "获取文档失败", err)
	}
	datasource, err := monitor.GetDatasource(ctx, knowledgeBase.DatasourceId)
	if err != nil {
		return nil, 0, global.NewError(500, "获取文档失败", err)
	}
	store, err := repository.NewVectorStore(datasource)
	if err != nil {
		return nil, 0, global.NewError(500, "获取文档失败", err)
	}
	return store.ListSlices(ctx, KnowledgeBaseCollectionName(knowledgeBase.Id), query)
}

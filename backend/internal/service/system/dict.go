package system

import (
	"context"
	"flowing/global"
	"flowing/internal/model/system"
)

func CreateDict(ctx context.Context, req system.CreateDictReq) error {
	dict := system.Dict{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}
	if err := system.CreateDict(ctx, dict); err != nil {
		return global.NewError(500, "新增字典失败", err)
	}
	return nil
}

func CreateDictItem(ctx context.Context, req system.CreateDictItemReq) error {
	dictItem := system.DictItem{
		DictId:      req.DictId,
		ItemKey:     req.ItemKey,
		ItemValue:   req.ItemValue,
		Description: req.Description,
		Sort:        req.Sort,
		Enable:      req.Enable,
	}
	if err := system.CreateDictItem(ctx, dictItem); err != nil {
		return global.NewError(500, "新增字典值失败", err)
	}
	return nil
}

func ListDictItems(ctx context.Context, req system.DictItemQuery) ([]system.DictItem, int64, error) {
	list, total, err := system.ListDictItem(ctx, req)
	if err != nil {
		return nil, 0, global.NewError(500, "查询字典值失败", err)
	}
	return list, total, nil
}

func ListDict(ctx context.Context, req system.DictQuery) ([]system.Dict, int64, error) {
	list, total, err := system.ListDict(ctx, req)
	if err != nil {
		return nil, 0, global.NewError(500, "查询字典失败", err)
	}
	return list, total, nil
}

func DeleteDict(ctx context.Context, id int64) error {
	if err := system.DeleteDict(ctx, id); err != nil {
		return global.NewError(500, "删除字典失败", err)
	}
	return nil
}

func ListDictItemByCode(ctx context.Context, code string) ([]system.DictItem, error) {
	list, err := system.ListDictItemByCode(ctx, code)
	if err != nil {
		return nil, global.NewError(500, "查询字典值失败", err)
	}
	return list, nil
}

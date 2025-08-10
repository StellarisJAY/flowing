package system

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
	"gorm.io/gorm"
)

type Dict struct {
	common.BaseModel
	Name        string `json:"name" gorm:"column:name;not null"`
	Code        string `json:"code" gorm:"column:code;type:varchar(255);unique;not null"`
	Description string `json:"description" gorm:"column:description;"`
}

func (Dict) TableName() string {
	return "sys_dict"
}

type DictItem struct {
	common.BaseModel
	DictId      int64  `json:"dictId,string" gorm:"column:dict_id;uniqueIndex:idx_dict_key;not null"`
	ItemKey     string `json:"itemKey" gorm:"column:item_key;type:varchar(255);uniqueIndex:idx_dict_key;not null"`
	ItemValue   string `json:"itemValue" gorm:"column:item_value;type:varchar(255);not null"`
	Description string `json:"description" gorm:"column:description;type:varchar(255);"`
	Sort        int    `json:"sort" gorm:"column:sort;default:1;"`
	Enable      bool   `json:"enable" gorm:"column:enable;"`
}

func (DictItem) TableName() string {
	return "sys_dict_item"
}

type DictQuery struct {
	common.BaseQuery
	Name string `json:"name" form:"name"`
	Code string `json:"code" form:"code"`
}

type DictItemQuery struct {
	common.BaseQuery
	DictId    int64  `json:"dictId" form:"dictId" binding:"required"`
	ItemKey   string `json:"itemKey" form:"itemKey"`
	ItemValue string `json:"itemValue" form:"itemValue"`
	Enable    *bool  `json:"enable" form:"enable"`
}

type CreateDictReq struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateDictItemReq struct {
	DictId      int64  `json:"dictId,string" binding:"required"`
	ItemKey     string `json:"itemKey" binding:"required"`
	ItemValue   string `json:"itemValue" binding:"required"`
	Description string `json:"description" binding:"required"`
	Sort        int    `json:"sort" binding:"required"`
	Enable      bool   `json:"enable" binding:"required"`
}

type ListDictItemByCodeReq struct {
	Code string `form:"code" binding:"required"`
}

func CreateDict(ctx context.Context, dict Dict) error {
	return repository.DB().WithContext(ctx).Create(&dict).Error
}

func CreateDictItem(ctx context.Context, dictItem DictItem) error {
	return repository.DB().WithContext(ctx).Create(&dictItem).Error
}

func DeleteDictItem(ctx context.Context, dictItem DictItem) error {
	return repository.DB().WithContext(ctx).Delete(&dictItem).Error
}

func DeleteDict(ctx context.Context, id int64) error {
	return repository.Tx(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Model(&Dict{}).Where("id = ?", id).Delete(&Dict{}).Error
		if err != nil {
			return err
		}
		return tx.WithContext(ctx).Model(&DictItem{}).Where("dict_id = ?", id).Delete(&DictItem{}).Error
	})
}

func ListDict(ctx context.Context, query DictQuery) ([]Dict, int64, error) {
	var list []Dict
	var total int64
	d := repository.DB().WithContext(ctx).Model(&Dict{})
	if query.Name != "" {
		d = d.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Code != "" {
		d = d.Where("code = ?", query.Code)
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).Find(&list).Error
	return list, total, err
}

func ListDictItem(ctx context.Context, query DictItemQuery) ([]DictItem, int64, error) {
	var list []DictItem
	var total int64
	d := repository.DB().WithContext(ctx).Model(&DictItem{}).Where("dict_id = ?", query.DictId)
	if query.ItemKey != "" {
		d = d.Where("item_key LIKE ?", "%"+query.ItemKey+"%")
	}
	if query.ItemValue != "" {
		d = d.Where("item_value LIKE ?", "%"+query.ItemValue+"%")
	}
	if query.Enable != nil {
		d = d.Where("enable = ?", *query.Enable)
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).Find(&list).Error
	return list, total, err
}

func ListDictItemByCode(ctx context.Context, code string) ([]DictItem, error) {
	var list []DictItem
	err := repository.DB().WithContext(ctx).
		Model(&DictItem{}).
		Joins("INNER JOIN sys_dict sd on sd.id = sys_dict_item.dict_id").
		Where("sd.code = ?", code).
		Find(&list).
		Error
	return list, err
}

package monitor

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
)

type DatasourceType string

const (
	DatasourceTypeMysql      DatasourceType = "mysql"
	DatasourceTypePostgresql DatasourceType = "postgresql"
	DatasourceTypeMilvus     DatasourceType = "milvus"
)

type Datasource struct {
	common.BaseModel
	Name        string         `json:"name" gorm:"column:name;type:varchar(128);not null;"`               // 数据源名称
	Code        string         `json:"code" gorm:"column:code;type:varchar(128);not null;unique;"`        // 数据源编码 唯一
	Type        DatasourceType `json:"type" gorm:"column:type;type:varchar(32);not null;"`                // 数据源类型
	Username    *string        `json:"username" gorm:"column:username;type:varchar(128);not null;"`       // 数据源用户名
	Password    *string        `json:"password" gorm:"column:password;type:varchar(128);not null;"`       // 数据源密码
	Host        string         `json:"host" gorm:"column:host;type:varchar(255);not null"`                // 数据源主机
	Port        int            `json:"port" gorm:"column:port;type:int(11);not null;"`                    // 数据源端口
	Database    *string        `json:"database" gorm:"column:database;type:varchar(128);not null;"`       // 数据源数据库
	Description string         `json:"description" gorm:"column:description;type:varchar(255);not null;"` // 数据源描述
}

func (Datasource) TableName() string {
	return "monitor_datasource"
}

type DatasourceQuery struct {
	common.BaseQuery
	Name string         `json:"name" form:"name"`
	Code string         `json:"code" form:"code"`
	Type DatasourceType `json:"type" form:"type"`
	Host string         `json:"host" form:"host"`
}

type CreateDatasourceReq struct {
	Name        string         `json:"name" binding:"required"`
	Code        string         `json:"code" binding:"required"`
	Type        DatasourceType `json:"type" binding:"required"`
	Username    *string        `json:"username"`
	Password    *string        `json:"password"`
	Host        string         `json:"host" binding:"required"`
	Port        int            `json:"port" binding:"required"`
	Database    *string        `json:"database"`
	Description string         `json:"description" binding:"required"`
}

type UpdateDatasourceReq struct {
	ID          int64          `json:"id,string" binding:"required"`
	Name        string         `json:"name" binding:"required"`
	Type        DatasourceType `json:"type" binding:"required"`
	Username    *string        `json:"username"`
	Password    *string        `json:"password"`
	Host        string         `json:"host" binding:"required"`
	Port        int            `json:"port" binding:"required"`
	Database    *string        `json:"database"`
	Description string         `json:"description" binding:"required"`
}

type PingDatasourceReq struct {
	Type     DatasourceType `json:"type" binding:"required"`
	Host     string         `json:"host" binding:"required"`
	Port     int            `json:"port" binding:"required"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	Database string         `json:"database"`
}

func CreateDatasource(ctx context.Context, model Datasource) error {
	return repository.DB(ctx).Create(&model).Error
}

func UpdateDatasource(ctx context.Context, model Datasource) error {
	return repository.DB(ctx).Model(&model).Updates(&model).Error
}

func ListDatasource(ctx context.Context, query DatasourceQuery) ([]Datasource, int64, error) {
	var list []Datasource
	var total int64
	d := repository.DB(ctx).Model(&Datasource{})
	if query.Name != "" {
		d = d.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Code != "" {
		d = d.Where("code LIKE ?", "%"+query.Code+"%")
	}
	if query.Type != "" {
		d = d.Where("type = ?", query.Type)
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if query.Host != "" {
		d = d.Where("host LIKE ?", "%"+query.Host+"%")
	}
	if err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func GetDatasource(ctx context.Context, id int64) (*Datasource, error) {
	var model Datasource
	if err := repository.DB(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

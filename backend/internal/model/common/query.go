package common

type BaseQuery struct {
	Page     bool `json:"page" form:"page"`
	PageNum  int  `json:"pageNum" form:"pageNum"`
	PageSize int  `json:"pageSize" form:"pageSize"`
}

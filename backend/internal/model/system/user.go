package system

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
)

type User struct {
	common.BaseModel
	Username string `json:"username" gorm:"column:username;type:varchar(50);unique;not null;"`
	NickName string `json:"nickName" gorm:"column:nick_name;type:varchar(50);not null;"`
	Password string `json:"password" gorm:"column:password;type:varchar(255);not null;"`
	Email    string `json:"email" gorm:"column:email;type:varchar(255);not null;"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(255);not null;"`
	Status   int    `json:"status" gorm:"column:status;type:int;default:1;"`

	Roles []Role `json:"roles" gorm:"many2many:sys_user_role;"`
}

func (User) TableName() string {
	return "sys_user"
}

type UserQuery struct {
	common.BaseQuery
	Username string `json:"username" form:"username"`
	NickName string `json:"nickName" form:"nickName"`
	Status   int    `json:"status" form:"status"`
}

func CreateUser(ctx context.Context, user *User) error {
	return repository.DB().WithContext(ctx).Create(user).Error
}

func GetUser(ctx context.Context, id int) (*User, error) {
	var user User
	err := repository.DB().WithContext(ctx).Where("id = ?", id).Preload("roles").First(&user).Error
	return &user, err
}

func ListUser(ctx context.Context, query UserQuery) ([]User, int64, error) {
	var users []User
	var total int64
	d := repository.DB().WithContext(ctx).Model(&User{})
	if query.Username != "" {
		d = d.Where("username LIKE ?", "%"+query.Username+"%")
	}
	if query.NickName != "" {
		d = d.Where("nick_name LIKE?", "%"+query.NickName+"%")
	}
	if query.Status != 0 {
		d = d.Where("status = ?", query.Status)
	}
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize, &total)).Preload("roles").Find(&users).Error
	return users, total, err
}

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
	Password string `json:"-" gorm:"column:password;type:varchar(255);not null;"`
	Email    string `json:"email" gorm:"column:email;type:varchar(255);not null;"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(255);not null;"`
	Status   int    `json:"status" gorm:"column:status;type:int;default:1;"`

	Roles []Role `json:"roles" gorm:"many2many:sys_user_role;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:role_id"`
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

type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type LoginReq struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Captcha    string `json:"captcha" binding:"required"`
	CaptchaKey string `json:"captchaKey" binding:"required"`
}

func CreateUser(ctx context.Context, user *User) error {
	return repository.DB().WithContext(ctx).Create(user).Error
}

func GetUser(ctx context.Context, username string) (*User, error) {
	var user User
	err := repository.DB().WithContext(ctx).Model(&User{}).Where("username = ?", username).First(&user).Preload("roles").Error
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
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize, &total)).Preload("roles").Scan(&users).Error
	return users, total, err
}

func CheckLogin(ctx context.Context, username string, password string) (bool, error) {
	var count int64
	err := repository.DB().WithContext(ctx).Model(&User{}).
		Where("username = ?", username).
		Where("password = ?", password).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

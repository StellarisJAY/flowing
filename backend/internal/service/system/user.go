package system

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flowing/global"
	sysmodel "flowing/internal/model/system"
	"flowing/internal/repository"
	"flowing/internal/util"
	"log/slog"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, user sysmodel.CreateUserReq) error {
	// 密码加密
	enc := sha256.New()
	password := hex.EncodeToString(enc.Sum([]byte(user.Password)))
	userModel := sysmodel.User{
		Username: user.Username,
		NickName: user.NickName,
		Password: password,
		Email:    user.Email,
		Phone:    user.Phone,
		Status:   1,
	}
	return repository.Tx(ctx, func(c context.Context) error {
		// 创建用户实体
		if err := sysmodel.CreateUser(c, &userModel); err != nil {
			return global.NewError(500, "创建用户失败", err)
		}
		// 创建用户角色关联
		if err := saveUserRole(c, userModel.Id, user.RoleIds); err != nil {
			return err
		}
		return nil
	})
}

func UpdateUser(ctx context.Context, user sysmodel.UpdateUserReq) error {
	userModel := sysmodel.User{
		NickName: user.NickName,
		Email:    user.Email,
		Phone:    user.Phone,
		Status:   user.Status,
	}
	return repository.Tx(ctx, func(c context.Context) error {
		// 更新用户实体
		if err := sysmodel.UpdateUser(c, userModel); err != nil {
			return global.NewError(500, "更新用户失败", err)
		}
		// 创建用户角色关联
		if err := saveUserRole(c, userModel.Id, user.RoleIds); err != nil {
			return err
		}
		return nil
	})
}

func saveUserRole(ctx context.Context, userId int64, roleIds []string) error {
	// 创建用户角色关联
	userRoles := make([]sysmodel.UserRole, len(roleIds))
	for i, roleId := range roleIds {
		roleId, err := strconv.ParseInt(roleId, 10, 64)
		if err != nil {
			continue
		}
		userRoles[i] = sysmodel.UserRole{
			UserId: userId,
			RoleId: roleId,
		}
	}
	// 删除所有用户角色关联
	err := repository.DB(ctx).Model(&sysmodel.UserRole{}).Where("user_id = ?", userId).Delete(&sysmodel.UserRole{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return global.NewError(500, "设置用户角色失败", err)
	}
	// 批量创建用户角色关联
	err = repository.DB(ctx).Model(&sysmodel.UserRole{}).CreateInBatches(userRoles, 10).Error
	if err != nil {
		return global.NewError(500, "设置用户角色失败", err)
	}
	return nil
}

func ListUser(ctx context.Context, query sysmodel.UserQuery) ([]sysmodel.User, int64, error) {
	return sysmodel.ListUser(ctx, query)
}

func GenCaptcha(_ context.Context) (string, string, error) {
	id, img, err := util.GenCaptcha()
	if err != nil {
		return "", "", global.NewError(500, "生成验证码失败", err)
	}
	return id, img, nil
}

func Login(ctx context.Context, req sysmodel.LoginReq) (string, error) {
	if !util.VerifyCaptcha(req.CaptchaKey, req.Captcha) {
		return "", global.NewError(400, "验证码错误", nil)
	}
	enc := sha256.New()
	password := hex.EncodeToString(enc.Sum([]byte(req.Password)))
	if ok, _ := sysmodel.CheckLogin(ctx, req.Username, password); !ok {
		slog.Info("用户登录失败", "username", req.Username)
		return "", global.NewError(400, "用户名或密码错误", nil)
	}
	// 获取用户信息
	user, err := sysmodel.GetUser(ctx, req.Username)
	if err != nil {
		return "", global.NewError(500, "登录失败", err)
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "flowing",
		Subject:   "access_token",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        uuid.New().String(),
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(repository.Config().Jwt.Secret))
	if err != nil {
		return "", global.NewError(500, "登录失败", err)
	}
	// 缓存用户信息
	userInfo, _ := json.Marshal(user)
	if err := repository.Redis().SetEx(ctx, claims.ID, string(userInfo), time.Hour*24).Err(); err != nil {
		return "", global.NewError(500, "登录失败", err)
	}
	return tokenString, nil
}

func GetUserMenus(ctx context.Context, userId int64) ([]*sysmodel.Menu, error) {
	menus, err := sysmodel.GetUserMenus(ctx, userId)
	if err != nil {
		return nil, global.NewError(500, "获取菜单失败", err)
	}
	return buildMenuTree(menus, false), nil
}

func DeleteUser(ctx context.Context, userId int64) error {
	return repository.Tx(ctx, func(c context.Context) error {
		// 删除用户角色关联
		if err := repository.DB(c).Delete(&sysmodel.UserRole{}, "user_id = ?", userId).Error; err != nil {
			return global.NewError(500, "删除用户失败", err)
		}
		// 删除用户
		if err := repository.DB(c).Delete(&sysmodel.User{}, "id = ?", userId).Error; err != nil {
			return global.NewError(500, "删除用户失败", err)
		}
		// TODO 其他用户关联表
		return nil
	})
}

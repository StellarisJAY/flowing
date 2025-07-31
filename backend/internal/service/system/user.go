package system

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flowing/global"
	sysmodel "flowing/internal/model/system"
	"flowing/internal/repository"
	"flowing/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func CreateUser(ctx context.Context, user sysmodel.CreateUserReq) error {
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
	return sysmodel.CreateUser(ctx, &userModel)
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

package system

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flowing/global"
	sysmodel "flowing/internal/model/system"
	"flowing/internal/repository"
	"log/slog"
)

const captchaKeyPrefix = "captcha_"

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

func GenCaptcha() (string, string, error) {
	// TODO 生成验证码
	return "todo", "todo", nil
}

func Login(ctx context.Context, req sysmodel.LoginReq) (string, error) {
	captcha, err := repository.Redis().Get(ctx, captchaKeyPrefix+req.CaptchaKey).Result()
	if err != nil || captcha != req.Captcha {
		slog.Info("验证码错误", "username", req.Username, "captcha", req.Captcha, "captchaKey", req.CaptchaKey)
		return "", global.NewError(400, "验证码错误", err)
	}
	enc := sha256.New()
	password := hex.EncodeToString(enc.Sum([]byte(req.Password)))
	if ok, _ := sysmodel.CheckLogin(ctx, req.Username, password); !ok {
		slog.Info("用户登录失败", "username", req.Username)
		return "", global.NewError(400, "用户名或密码错误", err)
	}
	// TODO 生成token
	return "token", nil
}

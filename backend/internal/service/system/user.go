package system

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	sysmodel "flowing/internal/model/system"
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

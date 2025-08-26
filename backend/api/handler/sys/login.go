package sys

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var login model.LoginReq
	if err := c.ShouldBindJSON(&login); err != nil {
		panic(global.ErrBadRequest(err))
	}
	resp, err := service.Login(c, login)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(resp))
}

func Logout(c *gin.Context) {
	cl, ok := c.Get("claims")
	if !ok {
		panic(global.ErrUnauthorized)
	}
	claims, ok := cl.(*jwt.RegisteredClaims)
	if !ok {
		panic(global.ErrUnauthorized)
	}
	if err := service.Logout(c, claims); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func GetCaptcha(c *gin.Context) {
	key, img, err := service.GenCaptcha(c)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(gin.H{"key": key, "captcha": img}))
}

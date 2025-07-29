package sys

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login model.LoginReq
	if err := c.ShouldBindJSON(&login); err != nil {
		panic(global.ErrBadRequest)
	}
	token, err := service.Login(c, login)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(token))
}

func Logout(c *gin.Context) {

}

func GetCaptcha(c *gin.Context) {
	key, img, err := service.GenCaptcha(c)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(gin.H{"key": key, "captcha": img}))
}

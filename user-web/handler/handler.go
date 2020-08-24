package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"net/http"

	user "liaotian/user-service/proto/user"

	"github.com/micro/go-micro/v2/client"
)

var (
	rpcUser user.UserService
)

func Init() {
	rpcUser = user.NewUserService("user.service.user", client.DefaultClient)
}

func Login(c *gin.Context) {

	var data user.Request
	_ = c.ShouldBindJSON(&data)

	res, err := rpcUser.Get(context.TODO(), &data)

	resultCode := http.StatusOK
	resultData := gin.H{}

	if err != nil {
		logger.Error(err)
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": res.Message,
		}
		return
	}

	if res.Code != 200 {
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": res.Message,
		}
		return
	}

	if res.User.Password == data.Password && res.User.Name == data.Name {
		resultData = gin.H{
			"message": res.Message,
			"user":    res.User,
		}
	} else {
		resultCode = http.StatusUnauthorized
		resultData = gin.H{
			"message": "用户名密码错误",
		}
	}

	c.JSON(resultCode, resultData)
	return
}

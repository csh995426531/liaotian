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

	resultData = gin.H{
		"status":  "posted",
		"message": res.Message,
		"user":    res.User,
	}
	c.JSON(resultCode, resultData)
	return
}

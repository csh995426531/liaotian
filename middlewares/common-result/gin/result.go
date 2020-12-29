package gin

import (
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, code int, data interface{}) {

	ctx.JSON(code, gin.H{
		"msg": "成功",
		"data": data,
	})
}

func Failed(ctx *gin.Context, code int, msg interface{}) {

	ctx.JSON(code, gin.H{
		"msg": msg,
		"data": nil,
	})
}
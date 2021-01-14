package gin

import (
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, code int, data interface{}) {

	ctx.JSON(code, gin.H{
		"msg":  "success",
		"data": data,
	})
}

func Failed(ctx *gin.Context, code int, msg interface{}) {

	var data interface{}
	ctx.JSON(code, gin.H{
		"msg":  msg,
		"data": data,
	})
}

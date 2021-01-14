package gin

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	authService "liaotian/domain/auth/proto"
	ginResult "liaotian/middlewares/common-result/gin"
	"liaotian/middlewares/logger/zap"
	"net/http"
)

/**
认证中间件
*/
func AuthMiddleware(domainAuth *authService.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		reqData := make(map[string]interface{})
		err := json.Unmarshal(data, &reqData)

		if err != nil || reqData["Token"] == nil || reqData["Token"].(string) == "" {
			ctx.Abort()
			ginResult.Failed(ctx, http.StatusUnauthorized, "请登录后操作")
			return
		}

		authReq := authService.ParseRequest{
			Token: reqData["Token"].(string),
		}

		res, err := (*domainAuth).Parse(ctx.Request.Context(), &authReq)
		if err != nil {
			ctx.Abort()
			zap.SugarLogger.Errorf("domainAuth.Parse error: %+v", err)
			ginResult.Failed(ctx, http.StatusUnauthorized, "请登录后操作")
			return
		}
		reqData["user_id"] = res.Data.UserId
		reqData["user_name"] = res.Data.Name
		newByte, _ := json.Marshal(reqData)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(newByte)) // 关键点
	}
}

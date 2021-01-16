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
	"strconv"
)

/**
认证中间件
*/
func AuthMiddleware(domainAuth *authService.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		reqData := make(map[string]interface{})
		_ = json.Unmarshal(data, &reqData)

		var token string
		if reqData["Token"] == nil || reqData["Token"].(string) == "" {

			token := ctx.GetHeader("token")
			if token == "" {
				ctx.Abort()
				ginResult.Failed(ctx, http.StatusUnauthorized, "请登录后操作")
				return
			}
		} else {
			token = reqData["Token"].(string)
		}

		authReq := authService.ParseRequest{
			Token: token,
		}

		res, err := (*domainAuth).Parse(ctx.Request.Context(), &authReq)
		if err != nil {
			ctx.Abort()
			zap.SugarLogger.Errorf("domainAuth.Parse error: %+v", err)
			ginResult.Failed(ctx, http.StatusUnauthorized, "请登录后操作")
			return
		}
		if ctx.Request.Method == "POST" {
			reqData["user_id"] = res.Data.UserId
			reqData["user_name"] = res.Data.Name
			newByte, _ := json.Marshal(reqData)
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(newByte))
		} else {
			values := ctx.Request.URL.Query()
			values.Add("user_id", strconv.FormatInt(res.Data.UserId, 10))
			values.Add("user_name", res.Data.Name)
			ctx.Request.URL.RawQuery = values.Encode()
		}
	}
}

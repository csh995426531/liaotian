package gin2micro

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/propagation"
	language_agent "github.com/SkyAPM/go2sky/reporter/grpc/language-agent"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/metadata"
)

const (
	componentIDGINHttpServer = 5006
	componentIDGoMicroClient = 5008
)

type routeInfo struct {
	operationName string
}

type middleware struct {
	routeMap     map[string]map[string]routeInfo
	routeMapOnce sync.Once
}

//Middleware gin middleware return HandlerFunc  with tracing.
func Middleware(engine *gin.Engine, tracer *go2sky.Tracer) gin.HandlerFunc {
	if engine == nil || tracer == nil {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	m := new(middleware)

	return func(c *gin.Context) {
		m.routeMapOnce.Do(func() {
			routes := engine.Routes()
			rm := make(map[string]map[string]routeInfo)
			for _, r := range routes {
				mm := rm[r.Method]
				if mm == nil {
					mm = make(map[string]routeInfo)
					rm[r.Method] = mm
				}
				mm[r.Handler] = routeInfo{
					operationName: fmt.Sprintf("/%s%s", r.Method, r.Path),
				}
			}
			m.routeMap = rm
		})
		var operationName string
		handlerName := c.HandlerName()
		if routeInfo, ok := m.routeMap[c.Request.Method][handlerName]; ok {
			operationName = routeInfo.operationName
		}
		if operationName == "" {
			operationName = c.Request.Method
		}

		span, ctx, err := tracer.CreateEntrySpan(c.Request.Context(), operationName, func() (string, error) {
			return c.Request.Header.Get(propagation.Header), nil
		})
		if err != nil {
			c.Next()
			return
		}

		span.SetComponent(componentIDGINHttpServer)
		span.Tag(go2sky.TagHTTPMethod, c.Request.Method)
		span.Tag(go2sky.TagURL, c.Request.Host+c.Request.URL.Path)
		span.SetSpanLayer(language_agent.SpanLayer_Http)

		/* 增加exitSpan方法，并把span id放入request.context内，以便go micro取出使用 */
		span2, err := tracer.CreateExitSpan(ctx, operationName, "gin-micro", func(header string) error {
			mda := map[string]string{propagation.Header: header, "user_name": "cuisaihang"}
			c.Request = c.Request.WithContext(metadata.MergeContext(ctx, mda, true))
			return nil
		})

		span2.SetComponent(componentIDGoMicroClient)
		span2.SetSpanLayer(language_agent.SpanLayer_RPCFramework)

		c.Next()

		if len(c.Errors) > 0 {
			span.Error(time.Now(), c.Errors.String())
		}
		span.Tag(go2sky.TagStatusCode, strconv.Itoa(c.Writer.Status()))
		span.End()
		span2.End()
	}
}

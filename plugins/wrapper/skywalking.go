package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
)

func NewLogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		//log.Printf("[Log Wrapper] Before serving request method: %v", req.Endpoint())
		err := fn(ctx, req, rsp)
		//log.Printf("[Log Wrapper] After serving request")
		return err
	}
}
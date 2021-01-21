package handler

import (
	"github.com/micro/go-micro/errors"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("user", "参数错误", http.StatusBadRequest)
)

func ErrorInternalServerError(err interface{}) error {
	return errors.InternalServerError("user", "内部错误，msg：%v", err)
}

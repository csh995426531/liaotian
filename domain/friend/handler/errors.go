package handler

import (
	"github.com/micro/go-micro/errors"
	"net/http"
)

var (
	ErrorBadRequest          = errors.New("friend", "参数错误", http.StatusBadRequest)
	ErrorApplicationNotFound = errors.New("friend", "申请单不存在", http.StatusNotFound)
)

func ErrorInternalServerError(err interface{}) error {
	return errors.InternalServerError("user", "内部错误，msg：%v", err)
}

package handler

import (
	"github.com/micro/go-micro/errors"
	"net/http"
)

var (
	ErrorBadRequest        = errors.New("user", "参数错误", http.StatusBadRequest)
	ErrorUserNotFound      = errors.New("user", "用户不存在", http.StatusNotFound)
	ErrorUserExists        = errors.New("user", "用户已存在", http.StatusForbidden)
	ErrorUserPasswordError = errors.New("user", "密码错误", http.StatusUnauthorized)
)

func ErrorInternalServerError(err interface{}) error {
	return errors.InternalServerError("user", "内部错误，msg：%v", err)
}

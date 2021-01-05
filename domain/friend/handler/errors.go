package handler

import "github.com/micro/go-micro/errors"

var (

)

func ErrorInternalServerError (err interface{}) error {
	return errors.InternalServerError("user", "内部错误，msg：%v", err)
}

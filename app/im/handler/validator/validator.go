package validator

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

//校验绑定参数
func Bind(ctx *gin.Context, reqValidator interface{}, req interface{}) (err error) {
	if ctx.Request.Method == "POST" {
		err = ctx.ShouldBind(reqValidator)
	} else {
		err = ctx.ShouldBindQuery(reqValidator)
	}
	if err != nil {
		return
	}

	uni := ut.New(en.New(), zh.New())
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()
	_ = zhtrans.RegisterDefaultTranslations(validate, trans)

	err = validate.Struct(reqValidator)

	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, msg := range errs.Translate(trans) {
				err = errors.New(fmt.Sprintf("%s", msg))
				return
			}
		}
		//err = translate.ErrTrans(err, translate.Trans)
		return
	}

	//验证通过，利用反射赋值到req
	reqValidatorVal := reflect.ValueOf(reqValidator).Elem()
	for i := 0; i < reqValidatorVal.NumField(); i++ {
		fieldName := reqValidatorVal.Type().Field(i).Name
		fieldValue := reqValidatorVal.Field(i).Interface()

		if reflect.ValueOf(req).Elem().FieldByName(fieldName).IsValid() {
			reflect.ValueOf(req).Elem().FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
		}
	}

	return
}

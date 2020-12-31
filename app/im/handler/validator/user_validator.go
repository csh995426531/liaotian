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
	userService "liaotian/domain/user/proto"
	//"liaotian/middlewares/validate/translate"
	"reflect"
)

//校验绑定参数
func Bind (ctx *gin.Context, reqValidator interface{}) (req userService.Request, err error) {

	err = ctx.ShouldBindJSON(reqValidator)
	if err != nil {
		err = errors.New("参数错误")
		return
	}

	en := en.New()
	zh := zh.New()
	uni := ut.New(en, zh)
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

		if reflect.ValueOf(&req).Elem().FieldByName(fieldName).IsValid() {
			reflect.ValueOf(&req).Elem().FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
		}
	}

	return
}

// 登录验证器
type LoginValidator struct {
	Account  string `validate:"required"`
	Password string `validate:"required"`
}

//注册验证器
type RegisterValidator struct {
	Account  string `validate:"required,min=1,max=20"`
	Name     string `validate:"required,min=1,max=20"`
	Password string `validate:"required,min=1,max=20"`
}

//获取用户信息验证器
type GetUserInfoValidator struct {
	Id  int64 `validate:"required"`
}

//更新用户信息验证器
type UpdateUserInfoValidator struct {
	Id    	 int64 `validate:"required"`
	Name 	 string
	Password string
	Avatar   string
}

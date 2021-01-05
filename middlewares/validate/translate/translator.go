package translate

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	//zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"liaotian/middlewares/logger/zap"
	//"reflect"
)

var (
	Trans *ut.Translator
)

func Init() {
	en := en.New()
	zh := zh.New()
	uni := ut.New(en, zh)
	trans, ok := uni.GetTranslator("zh")
	Trans = &trans
	if !ok {
		zap.SugarLogger.Error("全局验证器Init() 失败")
		return
	}
	zap.SugarLogger.Infof("全局验证器 初始化成功, %v", Trans)
}

//func With(validate *validator.Validate) (err error) {
//	fmt.Printf("validate2, type:%v, value:%+v\n",reflect.TypeOf(validate), validate)
//
//	//err = zhtrans.RegisterDefaultTranslations(validate, trans)
//	return
//}

func ErrTrans(err error, trans *ut.Translator) error {

	errs, ok := err.(validator.ValidationErrors)

	if ok {
		for field, err := range errs.Translate(*trans) {

			newErr := errors.New(fmt.Sprintf("%s哼哼%s", field, err))
			return newErr
		}
	}
	return err
}

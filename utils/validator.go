package utils

import (
	"errors"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

/*
Created by 斑斑砖 on 2023/8/14.
Description：
	参数校验
*/

var Validator = new(ValidateUtil)

type ValidateUtil struct{}

var (
	uni      *unTrans.UniversalTranslator
	validate *validator.Validate
)

func (v *ValidateUtil) Validate(data any) error {
	// 验证对象
	validate = validator.New()
	// 翻译器
	trans := v.validateTransInit(validate)

	err := validate.Struct(data)
	if err != nil {

		for _, e := range err.(validator.ValidationErrors) {

			return errors.New(e.Translate(trans))
		}
	}
	return nil
}

// 数据验证翻译器
func (*ValidateUtil) validateTransInit(validate *validator.Validate) unTrans.Translator {
	// 万能翻译器，保存所有的语言环境和翻译数据
	uni = unTrans.New(zh_Hans_CN.New())
	// 翻译器
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	//验证器注册翻译器
	zhTrans.RegisterDefaultTranslations(validate, trans)

	// 读取 Tag 中的 label 标签为字段的翻译
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	return trans
}

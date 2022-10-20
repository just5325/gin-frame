package validator

import (
	cnZH "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	ginValidator "github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// validate 实例
var ginValidate *ginValidator.Validate

// Validate/v10 全局验证器
var trans ut.Translator

func init() {
	// 实例化一个新的 validate 实例
	ginValidate = ginValidator.New()
	// 初始化Validate/v10国际化
	translations()
	// 添加自定义验证器
	registerValidation()
}

// 初始化Validate/v10国际化(返回中文错误信息)
func translations() {
	zh := cnZH.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	// 通过label标签返回自定义错误内容
	ginValidate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	err := zhTranslations.RegisterDefaultTranslations(ginValidate, trans)
	if err != nil {
		panic(err)
	}
}

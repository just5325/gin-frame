package validator

import (
	cnZH "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	ginValidator "github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var GinValidate *ginValidator.Validate

// Validate/v10 全局验证器
var trans ut.Translator

// 初始化Validate/v10国际化
func init() {
	zh := cnZH.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	GinValidate = ginValidator.New()

	//通过label标签返回自定义错误内容
	GinValidate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	zhTranslations.RegisterDefaultTranslations(GinValidate, trans)

	//自定义required_if错误内容
	GinValidate.RegisterTranslation("required_if", trans, func(ut ut.Translator) error {
		return ut.Add("required_if", "{0}为必填字段!", false) // see universal-translator for details
	}, func(ut ut.Translator, fe ginValidator.FieldError) string {
		t, _ := ut.T("required_if", fe.Field())
		return t
	})

}

// Translate 检验并返回检验错误信息
func Translate(err error) (errMsg string) {
	errs := err.(ginValidator.ValidationErrors)
	for _, err := range errs {
		errMsg = err.Translate(trans)
	}
	return
}

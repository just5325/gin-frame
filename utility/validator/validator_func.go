// 自定义全局验证器
// 注意!注意!注意!: 一些全局使用的自定义验证方法可以写在这里,如手机号格式验证,身份证号码格式验证等等...validator仅仅验证参数的格式即可.
//                带有业务逻辑的验证,就不建议写在这里了,
//                如验证指定订单号是否为指定用户的订单,
//                如验证数据库用户名是否存在,
//                这些应该显式的写在对应的业务逻辑中,比如service层封装了这个一个方法可以调用进行判断
// 创建人： 黄翠刚
// 创建时间： 2022.10.20

package validator

import (
	ut "github.com/go-playground/universal-translator"
	ginValidator "github.com/go-playground/validator/v10"
)

// 自定义验证器注册结构
type customFuncT struct {
	// 自定义验证器名称
	Tag string
	// 自定义验证器错误提示文字
	Text string
	// 自定义验证器方法
	Func ginValidator.Func
}

// 自定义验证器注册结构的切片类型
type customFuncListT []customFuncT

// 添加自定义验证器注册到全局
// 注意!注意!注意!：自定义验证器需要在这里注册,才能使用
var customFuncList = customFuncListT{
	// 添加一个自定义验证器
	{Tag: "demo", Text: "{0}必须为demo!", Func: demo},
}

// 添加自定义验证器
func registerValidation() {
	for _, customFunc := range customFuncList {
		err := ginValidate.RegisterValidation(customFunc.Tag, customFunc.Func)
		if err != nil {
			panic(err)
		}
		customFunc.registerTranslation(customFunc.Tag, customFunc.Text)
	}
}

// 注册自定义验证器的提示文字(也可以用这个方法自定义go-playground/validator包内置的验证方法的提示文字)
func (s *customFuncT) registerTranslation(tag string, text string) {
	//自定义required_if错误内容
	err := ginValidate.RegisterTranslation(tag, trans,
		func(ut ut.Translator) error {
			return ut.Add(tag, text, false)
		},
		func(ut ut.Translator, fe ginValidator.FieldError) string {
			t, _ := ut.T(tag, fe.Field())
			return t
		},
	)
	if err != nil {
		panic(err)
	}
}

// 一个示例自定义验证器,这里我们简单的验证参数的值 等于 "demo"吧
func demo(fl ginValidator.FieldLevel) (res bool) {
	if fl.Field().String() == "demo" {
		return true
	}
	return
}

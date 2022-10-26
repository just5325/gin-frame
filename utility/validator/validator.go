// 验证器验证方法
// 创建人： 黄翠刚
// 创建时间： 2022.10.19
// 使用示例:
// import utilityValidator "gin-frame/utility/validator"
// utilityValidator.GetInstance(ctx).Struct(&data)

package validator

import (
	"gin-frame/utility"
	"gin-frame/utility/response"
	"gin-frame/utility/response/response_code"
	"github.com/gin-gonic/gin"
	ginValidator "github.com/go-playground/validator/v10"
	"reflect"
)

// 声明结构体类型
type validatorImpl struct {
	ctx *gin.Context
}

// GetInstance 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func GetInstance(ctx *gin.Context) *validatorImpl {
	return &validatorImpl{
		ctx: ctx,
	}
}

// Translate 检验并返回检验错误信息
func (s *validatorImpl) Translate(err error) (errMsg string) {
	errs := err.(ginValidator.ValidationErrors)
	for _, err := range errs {
		errMsg = err.Translate(trans)
	}
	return
}

// Struct 结构体验证
func (s *validatorImpl) Struct(v interface{}) {
	utility.Common().ShouldBind(s.ctx, v)
	if err := ginValidate.Struct(v); err != nil {
		response.Response(s.ctx).Json(response_code.Error.Code, s.getValidMsg(err, v), nil)
		s.ctx.Abort()
	}
	return
}

// 获取验证错误信息
func (s *validatorImpl) getValidMsg(err error, obj interface{}) string {
	// 初始化一个变量,用于存储错误提示信息
	msg := ""
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(ginValidator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.StructField()); exist {
				// 获取用户自定义错误提示信息
				msg = f.Tag.Get("validateMsg")
				// 如果没有用户自定义错误提示信息,就使用验证器的错误信息
				if msg == "" {
					msg = s.Translate(err)
				}
				// 错误信息不需要全部返回，当找到第一个错误的信息时，就可以结束
				break
			}
		}
	} else {
		msg = err.Error()
	}

	return msg
}

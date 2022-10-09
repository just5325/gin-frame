// 返回数据工具包
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package response

import (
	"gin-frame/utility/response/response_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IResponse 声明接口类型
type IResponse interface {
	Json(code int, msg string, data interface{})
	SusJson(data interface{})
	FailJson(data interface{})
}

// 声明结构体类型
type responseImpl struct {
	ctx *gin.Context
}

// Response 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func Response(ctx *gin.Context) IResponse {
	return &responseImpl{
		ctx: ctx,
	}
}

// Json 返回json
func (s *responseImpl) Json(code int, msg string, data interface{}) {
	s.ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// SusJson 成功返回json
func (s *responseImpl) SusJson(data interface{}) {
	s.Json(response_code.Ok.Code, response_code.Ok.Message, data)
}

// FailJson 失败返回json
func (s *responseImpl) FailJson(data interface{}) {
	s.Json(response_code.Error.Code, response_code.Error.Message, data)
}

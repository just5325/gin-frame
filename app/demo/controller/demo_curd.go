package controller

import (
	"gin-frame/app/demo/validator"
	"gin-frame/utility/response"
	"github.com/gin-gonic/gin"
)

var DemoCurd = demoCurdController{}

type demoCurdController struct{}

// Create 数据库create操作
func (c *demoCurdController) Create(ctx *gin.Context) {
	// 参数验证,并获取绑定接口请求参数
	data := validator.DemoCurd.Create(ctx)
	response.Response(ctx).SusJson(data)
	return
}

package controller

import (
	"gin-frame/app/demo/service"
	"gin-frame/app/demo/validator"
	"gin-frame/utility"
	"gin-frame/utility/response"
	"gin-frame/utility/response/response_code"
	"github.com/gin-gonic/gin"
)

var DemoCurd = demoCurdController{}

type demoCurdController struct{}

// Create 数据库create操作
func (c *demoCurdController) Create(ctx *gin.Context) {
	// 获取绑定接口请求参数
	var data validator.DemoCurdValidatorCreate
	utility.Common().ShouldBind(ctx, &data)

	// 调用服务层方法完成操作
	res, err := service.DemoCurd(ctx).Create(data.Username, data.Password)
	if err != nil {
		response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
		return
	}
	response.Response(ctx).SusJson(res)
	return
}

// Update 数据库update操作
func (c *demoCurdController) Update(ctx *gin.Context) {
	// 获取绑定接口请求参数
	var data validator.DemoCurdValidatorUpdate
	utility.Common().ShouldBind(ctx, &data)

	// 调用服务层方法完成操作
	err := service.DemoCurd(ctx).Update(data.Username, data.Password)
	if err != nil {
		response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
		return
	}
	response.Response(ctx).SusJson(nil)
	return
}

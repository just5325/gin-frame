// demo控制器
// 仅仅是写个示例,从路由到控制器,控制器调用服务层完成并返回接口数据
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package controller

import (
	"gin-frame/app/demo/service"
	"gin-frame/utility/response"
	"gin-frame/utility/response/response_code"
	"github.com/gin-gonic/gin"
)

var Demo = demoController{}

type demoController struct{}

// Index 输出HelloWorld
func (c *demoController) Index(ctx *gin.Context) {
	res, err := service.Demo(ctx).HelloWorld()

	if err != nil {
		response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
		return
	}

	response.Response(ctx).SusJson(res)
	return
}

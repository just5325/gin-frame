// demo控制器
// 仅仅是写个示例,从路由到控制器,控制器调用服务层完成并返回接口数据
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package controller

import (
	"gin-frame/app/demo/service"
	"gin-frame/config"
	"gin-frame/utility/response"
	"gin-frame/utility/response/response_code"
	"github.com/gin-gonic/gin"
	"time"
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

// Shutdown 测试优雅关机
// 本方法会延迟8秒后返回结果,请求完本接口后,关闭本服务测试是否能请求完毕后优雅关机
func (c *demoController) Shutdown(ctx *gin.Context) {
	// 睡眠8秒
	time.Sleep(time.Duration(8) * time.Second)
	// 返回数据
	response.Response(ctx).SusJson(nil)
	return
}

// Config 读取配置文件
// 简单的读取一个配置文件即可
func (c *demoController) Config(ctx *gin.Context) {
	// 返回数据
	response.Response(ctx).SusJson(config.Config().GetViper().AllSettings())
	return
}

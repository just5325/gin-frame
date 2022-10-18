package router

import (
	"gin-frame/app/demo/controller"
	"github.com/gin-gonic/gin"
)

func demoRouter(group *gin.RouterGroup) {
	// 注册一个接口路由
	group.POST("/demo/index", controller.Demo.Index)
	// 测试优雅关机
	group.POST("/demo/shutdown", controller.Demo.Shutdown)
	// 读取配置文件
	group.POST("/demo/config", controller.Demo.Config)
	// 记录日志
	group.POST("/demo/log", controller.Demo.Log)
}

// 公共模块路由
// 如上传文件,发送短信,验证码等等公共接口,都可以写在公共模块下
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package router

import (
	"gin-frame/app/demo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(group *gin.RouterGroup) {
	// base路由分组
	baseRouterGroup := group.Group("/demo")
	// 注册一个接口路由
	baseRouterGroup.POST("/demo/index", controller.Demo.Index)
	// 测试优雅关机
	baseRouterGroup.POST("/demo/shutdown", controller.Demo.Shutdown)
	// 读取配置文件
	baseRouterGroup.POST("/demo/config", controller.Demo.Config)
	// 记录日志
	baseRouterGroup.POST("/demo/log", controller.Demo.Log)
}

package router

import "github.com/gin-gonic/gin"
import demoRouter "gin-frame/app/demo/router"

func InitRouter(r *gin.Engine) {
	// 绑定favicon.ico路由
	r.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.File("resource/image/favicon.ico")
	})

	// base路由分组
	baseRouterGroup := r.RouterGroup.Group("/api/v1")

	// 绑定 demo模块路由 到 base路由分组
	demoRouter.InitRouter(baseRouterGroup)
}

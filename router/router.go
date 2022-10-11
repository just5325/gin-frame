package router

import (
	"gin-frame/middleware"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)
import demoRouter "gin-frame/app/demo/router"

func InitRouter() (r *gin.Engine) {
	// 新建一个没有任何默认中间件的路由
	r = gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// Gin框架的请求ID中间件。使用X-Request-ID报头向响应添加标识符。如果X-Request-ID值在请求头中发送，则将其传递回调用方
	r.Use(requestid.New())

	// 绑定favicon.ico路由
	r.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.File("resource/image/favicon.ico")
	})

	// base路由分组
	baseRouterGroup := r.RouterGroup.Group(
		"/api/v1",
		// 中间件: api请求日志
		middleware.ApiLog().Handler,
	)

	// 绑定 demo模块路由 到 base路由分组
	demoRouter.InitRouter(baseRouterGroup)
	return
}

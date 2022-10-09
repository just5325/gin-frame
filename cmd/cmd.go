package cmd

import (
	"gin-frame/router"
	"github.com/gin-gonic/gin"
)

func Execute() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// 注册路由
	router.InitRouter(r)

	// 运行 http.Server
	err := r.Run()
	if err != nil {
		return
	}
}

package cmd

import (
	"context"
	"gin-frame/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func httpServer() {
	// 实例化路由
	r := router.InitRouter()

	// 下面代码来自于gin官网:https://gin-gonic.com/zh-cn/docs/examples/graceful-restart-or-stop/
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

package cmd

import (
	"context"
	"errors"
	"fmt"
	"gin-frame/config"
	"gin-frame/router"
	"net/http"
	"time"
)

func httpServer(stop <-chan interface{}, done chan<- error) {
	// 实例化路由
	r := router.InitRouter()

	// 定义http.Server
	s := &http.Server{
		Addr:    ":" + config.Config().GetViper().GetString("http_server.port"),
		Handler: r,
	}
	// 启动服务监听
	go func() {
		// 服务连接
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 向通道中写入错误信息
			done <- errors.New(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	// 接收stop，一旦通道中有数据了，或者close(stop)的话就会向下执行了,否则会一直阻塞在这里...
	<-stop
	// 优雅地关闭服务器（设置 10 秒的超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 关闭上下文环境中所有的协程
	if err := s.Shutdown(ctx); err != nil {
		// 向通道中写入错误信息
		done <- errors.New(fmt.Sprintf("Server Shutdown:%s\n", err))
	}
}

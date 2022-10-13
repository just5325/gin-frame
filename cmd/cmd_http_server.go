package cmd

import (
	"context"
	"fmt"
	"gin-frame/config"
	"gin-frame/router"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func httpServer(httpServerStop chan error, mainStop chan<- error) {
	// 实例化路由
	r := router.InitRouter()

	// 定义http.Server
	s := &http.Server{
		Addr:    ":" + config.Config().GetViper().GetString("http_server.port"),
		Handler: r,
	}

	// 监听系统中断信号以优雅地关闭服务器
	systemQuit(httpServerStop)

	go func() {
		// 接收stop，一旦通道中有数据了，或者close(stop)的话就会向下执行了,否则会一直阻塞在这里...
		<-httpServerStop
		// 优雅地关闭服务器（设置 10 秒的超时时间）
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// 关闭上下文环境中所有的协程
		_ = s.Shutdown(ctx)
		// 向通道中写入错误信息
		mainStop <- errors.New("server Shutdown")
	}()

	// 服务连接
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 向通道中写入错误信息
		mainStop <- errors.New(fmt.Sprintf("listen: %s\n", err))
	}
}

// 监听系统中断信号以优雅地关闭服务器
func systemQuit(httpServerStop chan error) {
	// 创建一个通道,监听系统中断信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		// 接收quit，没有数据，或者没有close的话会一直堵塞在这里
		<-quit
		// 结束子协程httpServer的生命周期
		httpServerStop <- errors.New("监听到系统中断信号,结束子协程httpServer的生命周期")
	}()
}

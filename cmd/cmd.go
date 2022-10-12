package cmd

import (
	"gin-frame/utility/log"
	"os"
	"os/signal"
)

// Execute 执行命令行命令
func Execute() {
	// 创建一个通道接收错误(有错误就会close(stop))
	done := make(chan error, 1)
	// 创建一个通道管理停止服务的信号(所有服务接收stop的数据后就可以停止服务了)
	stop := make(chan interface{})

	// 创建一个通道,等待系统中断信号以优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		// 接收quit，没有数据，或者没有close的话会一直堵塞在这里
		<-quit
		// 停止服务的信号
		stop <- true
	}()

	// 启动 http server
	go func() {
		httpServer(stop, done)
	}()

	if err := <-done; err != nil {
		log.Logger().GetLogger().Errorf("main err %+v", err)
		close(stop)
	}
}

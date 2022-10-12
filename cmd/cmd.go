package cmd

import (
	"gin-frame/utility/log"
)

// Execute 执行命令行命令
func Execute() {
	// 创建一个通道管理主协程的生命周期, mainStop接收到错误信息时,结束协程
	var mainStop = make(chan error, 1)
	// 创建一个通道管理子协程httpServer的生命周期, httpServerStop接收到错误消息时,结束协程
	var httpServerStop = make(chan error)

	// 启动 http server
	go func() {
		httpServer(httpServerStop, mainStop)
	}()

	// 等待mainStop写入错误信息, 未写入错误信息时阻塞主协程
	if err := <-mainStop; err != nil {
		log.Logger().GetLogger().Errorf("main err %+v", err)
		close(httpServerStop)
	}
}

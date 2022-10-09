// demo服务层
// 仅仅是写个示例,从路由到控制器,控制器调用服务层完成并返回接口数据
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package service

import (
	"github.com/gin-gonic/gin"
)

// IDemo 声明接口类型
type IDemo interface {
	HelloWorld() (string, error)
}

// 声明结构体类型
type demoImpl struct {
	ctx *gin.Context
}

// Demo 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func Demo(ctx *gin.Context) IDemo {
	return &demoImpl{
		ctx: ctx,
	}
}

// HelloWorld 方法
func (s *demoImpl) HelloWorld() (string, error) {
	return "你好,世界!", nil
}

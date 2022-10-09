// 工具包
// 一些比较公用的工具方法可以集中写在这里,方便使用
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package utility

import (
	"fmt"
)

// ICommon 声明接口类型
type ICommon interface {
	// FmtPrint 方便打印调试
	FmtPrint(data interface{})
}

// 声明结构体类型
type commonImpl struct{}

// Common 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func Common() ICommon {
	return &commonImpl{}
}

// FmtPrint 方便打印调试
func (s *commonImpl) FmtPrint(data interface{}) {
	fmt.Printf("打印开始----------------------------------------------------------------------- \n")
	fmt.Printf("%+v \n", data)
	fmt.Printf("打印结束----------------------------------------------------------------------- \n\n\n")
}

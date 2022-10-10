// 工具包
// 一些比较公用的工具方法可以集中写在这里,方便使用
// 创建人： 黄翠刚
// 创建时间： 2022.10.09
// 使用示例:utility.Common().FmtPrint()

package utility

import (
	"fmt"
	"os"
)

// ICommon 声明接口类型
type ICommon interface {
	// FmtPrint 方便打印调试
	FmtPrint(data interface{})
	// FileIsExisted 文件或文件夹是否存在
	FileIsExisted(filename string) bool
	// MkDir 新建文件夹
	MkDir(src string) error
	// IsNotExistMkDir 检查文件夹是否存在，如果不存在则新建文件夹
	IsNotExistMkDir(src string) error
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

// FileIsExisted 文件或文件夹是否存在
func (s *commonImpl) FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// MkDir 新建文件夹
func (s *commonImpl) MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// IsNotExistMkDir 检查文件夹是否存在，如果不存在则新建文件夹
func (s *commonImpl) IsNotExistMkDir(src string) error {
	if exist := s.FileIsExisted(src); exist == false {
		if err := s.MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

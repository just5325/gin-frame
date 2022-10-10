// 配置管理
// 基于github.com/spf13/viper扩展的单例模式,
// 使用方式:config.Config().GetViper() 获取 *viper.Viper,后续所有操作请直接参考github.com/spf13/viper
// 创建人： 黄翠刚
// 创建时间： 2022.10.10

package config

import (
	"errors"
	"fmt"
	"gin-frame/utility"
	"github.com/spf13/viper"
	"sync"
)

// 声明一个变量,用于存储 *viper.Viper
var configs *viper.Viper

// 配置文件路径
var configFile = "./config/config.yml"

// 包初始化()
func init() {

	if configs != nil {
		return
	}

	var once sync.Once
	once.Do(func() {
		configs = initConfigs()
	})
}

// 初始化配置信息
func initConfigs() (configs *viper.Viper) {
	// 检查配置文件是否存在
	if utility.Common().FileIsExisted(configFile) == false {
		err := errors.New(fmt.Sprintf("缺少配置文件:%s", configFile))
		panic(err)
	}

	// 导入配置文件
	viper.SetConfigFile(configFile)

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		// 读取配置文件出错
		panic(err)
	}

	return viper.GetViper()
}

// IConfig 声明接口类型
type IConfig interface {
	// GetViper 获取 *viper.Viper,后续所有操作请直接参考github.com/spf13/viper
	GetViper() *viper.Viper
}

// 声明结构体类型
type configImpl struct{}

// Config 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func Config() IConfig {
	return &configImpl{}
}

// GetViper 获取 *viper.Viper,后续所有操作请直接参考github.com/spf13/viper
func (s *configImpl) GetViper() *viper.Viper {
	return configs
}

// 配置管理
// 基于github.com/spf13/viper扩展的单例模式,
// 使用方式:config.GetInstance().GetViper() 获取 *viper.Viper,后续所有操作请直接参考github.com/spf13/viper
// 创建人： 黄翠刚
// 创建时间： 2022.10.10

package config

import (
	"github.com/spf13/viper"
)

// 声明结构体类型
type configImpl struct{}

// GetInstance 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func GetInstance() *configImpl {
	return &configImpl{}
}

// GetViper 获取 *viper.Viper,后续所有操作请直接参考github.com/spf13/viper
func (s *configImpl) GetViper() *viper.Viper {
	return configs
}

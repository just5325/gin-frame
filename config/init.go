package config

import (
	"gin-frame/utility"
	"github.com/spf13/viper"
	"os"
	"path"
	"sync"
)

// 声明一个变量,用于存储 *viper.Viper
var configs *viper.Viper

// 配置文件路径
type configFileT struct {
	filename string
	filepath string
}

var configFile = configFileT{
	filename: "config.yml",
	filepath: "config",
}

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
	// 获取当前程序的根目录(绝对路径)
	pwd, _ := os.Getwd()
	// 配置文件目录
	configFilePath := path.Join(pwd, configFile.filepath)
	// 配置文件地址
	configFileName := path.Join(pwd, configFile.filepath, configFile.filename)

	// 检查配置文件目录是否存在，如果不存在则新建文件夹
	if err := utility.Common().IsNotExistMkDir(configFilePath); err != nil {
		panic(err)
	}
	// 检查配置文件是否存在，如果不存在则新建文件
	if utility.Common().FileIsExisted(configFileName) == false {
		// 创建空配置文件(配置文件中所有的配置都有默认值,这里创建空配置文件即可)
		err := os.WriteFile(configFileName, []byte(""), 0666)
		if err != nil {
			panic(err)
		}
	}

	// 导入配置文件
	viper.SetConfigFile(configFileName)

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		// 读取配置文件出错
		panic(err)
	}

	return viper.GetViper()
}

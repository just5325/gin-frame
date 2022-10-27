package log

import (
	"bufio"
	"gin-frame/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
	"time"
)

// 声明一个变量,用于存储 *logrus.Logger
var logger *logrus.Logger

// 默认日志文件的软链接路径(程序根目录的相对路径)
var linkPath = "resource/logs/log.log"

// 默认日志切割生成的日志文件路径(程序根目录的相对路径)
var filePath = "resource/logs/log/%Y%m%d%H.log"

// 默认日志文件最大保存天数
var withMaxAge = 30

// 默认日志切割时间间隔(单位小时)(隔多久分割一次)
var withRotationTime = 1

// 包初始化
func init() {

	if logger != nil {
		return
	}

	var once sync.Once
	once.Do(func() {
		logger = logrus.New()
		initLogger()
	})
}

func initLogger() {
	// Logger.Out = os.Stdout
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(src)

	logger.Out = writer // config.GetLogConfig().FilePath

	//设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 获取当前程序的根目录(绝对路径)
	pwd, _ := os.Getwd()

	// 获取配置日志文件的软链接路径
	if config.GetInstance().GetViper().GetString("log.link_path") != "" {
		linkPath = path.Join(pwd, config.GetInstance().GetViper().GetString("log.link_path"))
	} else {
		linkPath = path.Join(pwd, linkPath)
	}

	// 获取配置日志切割生成的路径
	if config.GetInstance().GetViper().GetString("log.file_path") != "" {
		filePath = path.Join(pwd, config.GetInstance().GetViper().GetString("log.file_path"))
	} else {
		filePath = path.Join(pwd, filePath)
	}

	// 获取配置日志文件最大保存天数
	if config.GetInstance().GetViper().GetInt("log.with_max_age") > 0 {
		withMaxAge = config.GetInstance().GetViper().GetInt("log.with_max_age")
	}
	// 获取配置日志切割时间间隔
	if config.GetInstance().GetViper().GetInt("log.with_rotation_time") > 0 {
		withRotationTime = config.GetInstance().GetViper().GetInt("log.with_rotation_time")
	}

	logWriter, _ := rotatelogs.New(
		filePath,
		// 生成软链 指向最新的日志文件 - window 下使用要开启开发者模式
		rotatelogs.WithLinkName(linkPath),
		// 文件最大保存时间
		rotatelogs.WithMaxAge(time.Duration(withMaxAge*24)*time.Hour),
		// 设置日志切割时间间隔(1小时)(隔多久分割一次)
		rotatelogs.WithRotationTime(time.Duration(withRotationTime)*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
}

// 日志工具包
// 基于github.com/sirupsen/logrus
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package log

import (
	"bufio"
	"fmt"
	"gin-frame/config"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
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

// MyWriter 实现 logger 的interface{}
type MyWriter struct {
	mlog *logrus.Logger
}

func (m *MyWriter) Printf(format string, v ...interface{}) {
	logStr := fmt.Sprintf(format, v...)
	m.mlog.Info(logStr)
}

func NewMyWriter() *MyWriter {
	return &MyWriter{mlog: logger}
}

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
	if config.Config().GetViper().GetString("log.link_path") != "" {
		linkPath = path.Join(pwd, config.Config().GetViper().GetString("log.link_path"))
	} else {
		linkPath = path.Join(pwd, linkPath)
	}

	// 获取配置日志切割生成的路径
	if config.Config().GetViper().GetString("log.file_path") != "" {
		filePath = path.Join(pwd, config.Config().GetViper().GetString("log.file_path"))
	} else {
		filePath = path.Join(pwd, filePath)
	}

	// 获取配置日志文件最大保存天数
	if config.Config().GetViper().GetInt("log.with_max_age") > 0 {
		withMaxAge = config.Config().GetViper().GetInt("log.with_max_age")
	}
	// 获取配置日志切割时间间隔
	if config.Config().GetViper().GetInt("log.with_rotation_time") > 0 {
		withRotationTime = config.Config().GetViper().GetInt("log.with_rotation_time")
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

// ILogger 声明接口类型
type ILogger interface {
	// GetLogger 获取 *logrus.Logger,后续所有操作请直接参考github.com/sirupsen/logrus
	GetLogger() *logrus.Logger
	// Log 记录日志
	Log(ctx *gin.Context, fields logrus.Fields)
}

// 声明结构体类型
type loggerImpl struct{}

// Logger 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func Logger() ILogger {
	return &loggerImpl{}
}

// GetLogger 获取 *logrus.Logger,后续所有操作请直接参考github.com/sirupsen/logrus
func (s *loggerImpl) GetLogger() *logrus.Logger {
	return logger
}

// Log 记录日志
func (s *loggerImpl) Log(ctx *gin.Context, fields logrus.Fields) {
	// 请求ID (整个HTTP请求生命周期内,记录的日志 request_id 都是相同的)
	fields["request_id"] = requestid.Get(ctx)
	Logger().GetLogger().WithFields(fields).Info()
}

// 日志工具包
// 基于github.com/sirupsen/logrus
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package log

import (
	"fmt"
	"gin-frame/config"
	"gin-frame/utility"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"os"
	"sync"
	"time"
)

// 声明一个变量,用于存储 *logrus.Logger
var logger *logrus.Logger

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
	// 获取配置的日志文件目录
	logDir := config.Config().GetViper().GetString("log.log_dir")
	// 日志文件目录拼接年月作为二级目录
	logDir += fmt.Sprintf("/%d-%d", time.Now().Year(), cast.ToInt(time.Now().Format("01")))
	// 日志文件
	logFile := fmt.Sprintf("%s/%d.log", logDir, time.Now().Day())
	// 检查文件夹是否存在，如果不存在则新建文件夹
	err := utility.Common().IsNotExistMkDir(logDir)
	if err != nil {
		panic(err)
	}

	//设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
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

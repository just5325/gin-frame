// 日志工具包
// 基于github.com/sirupsen/logrus
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package log

import (
	"fmt"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogType struct {
	Type string
	Msg  string
}

// 内置的几种日志类型
var (
	// LOG 普通日志,比如手动在关键位置记录个日志什么的方便查询用的
	LOG = LogType{Type: "log", Msg: ""}
	// PANIC_LOG 程序发生panic记录的日志
	PANIC_LOG = LogType{Type: "panic_log", Msg: "panic日志"}
	// API_LOG api请求记录的日志,比如请求参数,返回参数什么什么的
	API_LOG = LogType{Type: "api_log", Msg: "api请求日志"}
	// SQL_LOG 数据库的SQL日志
	SQL_LOG = LogType{Type: "sql_log", Msg: "数据库SQL日志"}
)

// 声明结构体类型
type loggerImpl struct{}

// GetInstance 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func GetInstance() *loggerImpl {
	return &loggerImpl{}
}

// GetLogger 获取 *logrus.Logger,后续所有操作请直接参考github.com/sirupsen/logrus
func (s *loggerImpl) GetLogger() *logrus.Logger {
	return logger
}

// Log 记录日志
func (s *loggerImpl) Log(ctx *gin.Context, logType LogType, logFields logrus.Fields) {
	// 设置默认日志类型
	if logType.Type == "" {
		logType.Type = LOG.Type
	}
	// 记录日志的字段
	fields := logrus.Fields{
		// 请求ID (整个HTTP请求生命周期内,记录的日志 request_id 都是相同的)
		"request_id": requestid.Get(ctx),
		// 日志详细内容
		"info": logFields,
		// 日志类型
		"log_type": logType.Type,
	}
	s.GetLogger().WithFields(fields).Info(logType.Msg)
}

// Printf gorm的日志驱动需要实现Printf方法,进行记录数据库SQL日志
func (s *loggerImpl) Printf(_ string, v ...interface{}) {
	fields := logrus.Fields{
		// 日志类型
		"log_type": SQL_LOG.Type,
		// 日志详细内容
		"info": gin.H{
			// 执行SQL的文件路径
			"file": v[0],
			// 执行SQL消耗的时间(单位毫秒)
			"time": fmt.Sprintf("%.3fms", v[1]),
			// 执行SQL影响的行
			"rows": v[2],
			// SQL
			"sql": v[3],
		},
	}
	s.GetLogger().WithFields(fields).Info(SQL_LOG.Msg)
}

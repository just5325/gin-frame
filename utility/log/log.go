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
func (s *loggerImpl) Log(ctx *gin.Context, msg string, logFields logrus.Fields) {
	fields := logrus.Fields{
		// 请求ID (整个HTTP请求生命周期内,记录的日志 request_id 都是相同的)
		"request_id": requestid.Get(ctx),
		// 日志详细内容
		"info": logFields,
	}
	s.GetLogger().WithFields(fields).Info(msg)
}

// Printf gorm的日志驱动需要实现Printf方法,进行记录数据库SQL日志
func (s *loggerImpl) Printf(format string, v ...interface{}) {
	fields := logrus.Fields{
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
	s.GetLogger().WithFields(fields).Info("数据库SQL日志")
}

package custom_recovery

import (
	"fmt"
	"gin-frame/utility/log"
	"gin-frame/utility/response"
	"gin-frame/utility/response/response_code"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var CustomRecovery gin.RecoveryFunc = func(ctx *gin.Context, err any) {
	// 记录日志
	log.GetInstance().Log(ctx, log.PANIC_LOG, logrus.Fields{
		"err": fmt.Sprintf("%+v", err),
	})
	// 返回API友好报错信息
	response.Response(ctx).Json(response_code.Error.Code, "程序异常!", nil)
	// TODO 这里最好能写个消息推送的机制,程序中发生panic的时候即时推送错误信息到开发人员的邮箱中,方便及时处理问题 ...

	ctx.Abort()
	return
}

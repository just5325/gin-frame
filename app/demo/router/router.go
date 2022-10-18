// 公共模块路由
// 如上传文件,发送短信,验证码等等公共接口,都可以写在公共模块下
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(group *gin.RouterGroup) {
	// base路由分组
	baseRouterGroup := group.Group("/demo")
	// demo 控制器相关路由注册
	demoRouter(baseRouterGroup)
}

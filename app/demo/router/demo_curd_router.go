package router

import (
	"gin-frame/app/demo/controller"
	"gin-frame/app/demo/validator"
	"github.com/gin-gonic/gin"
)

// demo_curd 控制器相关路由注册
func demoCurdRouter(group *gin.RouterGroup) {
	// demo数据库create操作
	group.POST("/demo_curd/create", validator.DemoCurd.Create, controller.DemoCurd.Create)
}

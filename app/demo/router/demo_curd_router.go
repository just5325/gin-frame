package router

import (
	"gin-frame/app/demo/controller"
	appValidator "gin-frame/app/demo/validator"
	"github.com/gin-gonic/gin"
)

// demo_curd 控制器相关路由注册
func demoCurdRouter(group *gin.RouterGroup) {
	// demo数据库create操作
	group.POST("/demo_curd/create", appValidator.DemoCurd.Create, controller.DemoCurd.Create)
	// demo数据库update操作
	group.POST("/demo_curd/update", appValidator.DemoCurd.Update, controller.DemoCurd.Update)
}

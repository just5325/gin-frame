package router

import (
	"gin-frame/app/demo/controller"
	"github.com/gin-gonic/gin"
)

func redisRouter(group *gin.RouterGroup) {
	group.POST("/redis/example", controller.Redis.Example)
}

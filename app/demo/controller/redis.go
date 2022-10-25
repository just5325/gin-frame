package controller

import (
	"gin-frame/utility"
	"gin-frame/utility/redis"
	"gin-frame/utility/response"
	"gin-frame/utility/response/response_code"
	"github.com/gin-gonic/gin"
	goRedis "github.com/go-redis/redis/v9"
)

var Redis = redisController{}

type redisController struct{}

// Example redis示例操作
func (c *redisController) Example(ctx *gin.Context) {
	// redis set
	err := redis.GetRedisClient().Set(ctx, "key", "value", 0).Err()
	if err != nil {
		response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
		return
	}

	// redis get
	val, err := redis.GetRedisClient().Get(ctx, "key").Result()
	if err != nil {
		response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
		return
	}
	utility.Common().FmtPrint(val)

	val2, err := redis.GetRedisClient().Get(ctx, "key2").Result()
	if err == goRedis.Nil {
		utility.Common().FmtPrint("key2 不存在")
	} else if err != nil {
		response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
		return
	} else {
		utility.Common().FmtPrint(val2)
	}
}

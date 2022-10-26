package token

import (
	"crypto/md5"
	"fmt"
	"gin-frame/config"
	"gin-frame/utility/redis"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	goRedis "github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

// token 的前缀
var tokenPrefix = "token"

// token 有效期, 默认30天
var tokenExpiration = 30 * 24 * time.Hour

// Options 生成 token 的选项
type Options struct {
	// 用户ID
	Id int
	// 用户ID类型
	// 如:小程序登录的可以填:wxMini PC端登录的:pc 苹果手机登录的:ios 安卓手机登录的:android ipad登录的:ipad
	// 这个就可以实现一个用户ID可以同时持有多个不同端的有效token, 并且相同端只能有一个有效token
	// 如果不想实现这个多端token的功能的话, 不管什么端登录的都写死一个固定的值就行了,比如写死为"token"就行
	Type string
	// 用户token缓存数据
	TokenData string
}

// 声明结构体类型
type tokenImpl struct {
	ctx *gin.Context
}

// 初始化
func init() {
	// 获取配置文件的 token 有效期
	if config.Config().GetViper().GetInt("token.expire") > 0 {
		tokenExpiration = time.Duration(config.Config().GetViper().GetInt("token.expire")) * 24 * time.Hour
	}
}

// GetInstance 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func GetInstance(ctx *gin.Context) *tokenImpl {
	return &tokenImpl{
		ctx: ctx,
	}
}

// NewToken 生成一个新的 token
func (s *tokenImpl) NewToken(options Options) (token string, err error) {
	// 生成 [min,max)之间的随机数
	min := 10000000
	max := 99999999
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(max-min) + min

	// 生成 token (主要是生成一个唯一不重复的字符串, 然后MD5而已)
	token = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v-%v-%v-%v-%v", options.Id, options.Type, time.Now().UnixNano(), randInt, requestid.Get(s.ctx)))))

	// 唯一ID缓存key
	idKey := fmt.Sprintf("%v:%v", options.Type, options.Id)

	// 通过 唯一ID缓存key 获取指向的 token
	oldToken, err := redis.GetRedisClient().GetDel(s.ctx, idKey).Result()
	if err != nil && err != goRedis.Nil {
		return
	}
	// 删除原有 token,使之失效
	if len(oldToken) > 0 {
		redis.GetRedisClient().Del(s.ctx, s.getRedisTokenKey(oldToken))
	}

	// 设置  唯一ID缓存key 指向新的 token
	redis.GetRedisClient().Set(s.ctx, idKey, token, 0)

	// 设置 token 缓存的用户信息
	err = redis.GetRedisClient().Set(s.ctx, s.getRedisTokenKey(token), options.TokenData, tokenExpiration).Err()

	return
}

// ParseToken 解析token
func (s *tokenImpl) ParseToken(token string) (tokenData interface{}, err error) {
	// 获取 token 的缓存数据
	tokenData, err = redis.GetRedisClient().Get(s.ctx, s.getRedisTokenKey(token)).Result()
	if err == goRedis.Nil {
		err = errors.New("token无效")
	}
	return
}

// 获取redis缓存的token的key值(方便统一修改或配置前缀之类的东西)
func (s *tokenImpl) getRedisTokenKey(token string) string {
	return fmt.Sprintf("%s:%s", tokenPrefix, token)
}

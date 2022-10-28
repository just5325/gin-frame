// demo控制器
// 仅仅是写个示例,从路由到控制器,控制器调用服务层完成并返回接口数据
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package controller

import (
	"encoding/json"
	"gin-frame/app/demo/service"
	"gin-frame/config"
	"gin-frame/utility"
	"gin-frame/utility/log"
	"gin-frame/utility/response"
	"gin-frame/utility/response/response_code"
	"gin-frame/utility/token"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"time"
)

var Demo = demoController{}

type demoController struct{}

// Index 输出HelloWorld
func (c *demoController) Index(ctx *gin.Context) {
	res, err := service.Demo(ctx).HelloWorld()

	if err != nil {
		response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
		return
	}

	response.Response(ctx).SusJson(res)
	return
}

// Shutdown 测试优雅关机
// 本方法会延迟8秒后返回结果,请求完本接口后,关闭本服务测试是否能请求完毕后优雅关机
func (c *demoController) Shutdown(ctx *gin.Context) {
	// 睡眠8秒
	time.Sleep(time.Duration(8) * time.Second)
	// 返回数据
	response.Response(ctx).SusJson(nil)
	return
}

// Config 读取配置文件
// 简单的读取一个配置文件即可
func (c *demoController) Config(ctx *gin.Context) {
	// 返回数据
	response.Response(ctx).SusJson(config.GetInstance().GetViper().AllSettings())
	return
}

// Log 记录日志
func (c *demoController) Log(ctx *gin.Context) {
	// 仅仅演示一下怎么记录日志而已
	log.GetInstance().Log(ctx, log.LogType{Msg: "仅仅演示一下怎么记录日志而已"}, logrus.Fields{
		"a": 1,
		"b": 2,
	})
	return
}

// UserPermission 用户权限
type UserPermission struct {
	// 权限key(不同的权限key需要保持唯一)
	Permission string `json:"key"`
	// 权限名称
	Title string `json:"title"`
}

// TokenData 用户token缓存的数据接口
type TokenData struct {
	// 用户ID
	Uid int `json:"uid"`
	// 用户名
	UserName string `json:"username"`
	// 用户权限
	Permissions []UserPermission `json:"permissions"`
}

// Token 获取token
func (c *demoController) Token(ctx *gin.Context) {
	// 模拟登录接口, 登录成功后查询出来的用户权限
	permissions := make([]UserPermission, 0)
	permissions = append(permissions, UserPermission{Permission: "order/info", Title: "查看订单"})
	permissions = append(permissions, UserPermission{Permission: "order/create", Title: "创建订单"})
	permissions = append(permissions, UserPermission{Permission: "order/edit", Title: "编辑订单"})
	// 用户token缓存的数据接口
	tokenData := TokenData{
		Uid:         3,
		UserName:    "黄翠刚",
		Permissions: permissions,
	}
	// redis 中只能存json字符串, 所以这里把token缓存的数据转换成json字符串
	tokenDataByte, err := json.Marshal(tokenData)
	if err != nil {
		return
	}
	// 获取新的token
	res, err := token.GetInstance(ctx).NewToken(token.Options{Id: tokenData.Uid, Type: "admin:wap", TokenData: string(tokenDataByte)})
	if err != nil {
		response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
		return
	}
	response.Response(ctx).SusJson(gin.H{
		"token": res,
	})
	return
}

// 为路由 /api/v1/demo/demo/parseToken 设置的接口参数结构体
type demoParseToken struct {
	Token string `json:"token"`
}

// ParseToken 解析token
func (c *demoController) ParseToken(ctx *gin.Context) {
	// 获取绑定接口请求参数
	var data demoParseToken
	utility.Common().ShouldBind(ctx, &data)

	// 解析token
	tokenDataString, err := token.GetInstance(ctx).ParseToken(data.Token)
	if err != nil {
		response.Response(ctx).Json(response_code.TokenInvalid.Code, response_code.TokenInvalid.Message, nil)
		return
	}

	var responseData interface{}

	// 判断token中存储的数据是否为有效的json格式字符串
	jsonBlob := []byte(cast.ToString(tokenDataString))
	if json.Valid(jsonBlob) {
		var tokenData TokenData
		err = json.Unmarshal(jsonBlob, &tokenData)
		if err != nil {
			response.Response(ctx).Json(response_code.Error.Code, err.Error(), nil)
			return
		}
		responseData = tokenData
	} else {
		responseData = tokenDataString
	}

	response.Response(ctx).SusJson(responseData)
	return
}

// Recovery 模拟接口中发生panic的表现
func (c *demoController) Recovery(_ *gin.Context) {
	err := errors.New("模拟接口中发生panic的表现")
	panic(err)
	return
}

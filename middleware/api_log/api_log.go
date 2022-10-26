// 中间件: api请求日志
//
// 创建人： 黄翠刚
// 创建时间： 2022.10.10

package api_log

import (
	"bytes"
	"encoding/json"
	"gin-frame/utility"
	"gin-frame/utility/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 声明结构体类型
type apiLogImpl struct{}

// GetInstance 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func GetInstance() *apiLogImpl {
	return &apiLogImpl{}
}

// Handler 中间件逻辑执行体
func (s *apiLogImpl) Handler(ctx *gin.Context) {
	blw := &customResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw

	ctx.Next()

	// 获取所有请求参数
	inputs, err := utility.Common().RequestInputs(ctx)
	if err != nil {
		panic(err)
	}

	// 获取返回数据
	responseData := blw.body.String()
	var response interface{}
	// 检查数据是否为有效的json格式字符串,是的话转换为struct,避免记录日志后数据被加"\"反斜线,示例效果如下
	// {"request_id":"e0be6217-8f6f-4e0d-865a-530cc7f9fd4e","response":"{\"code\":200,\"data\":\"你好,世界!\",\"msg\":\"请求成功！\"}","time":"2022-10-11 09:32:30","url":"/api/v1/demo/demo/index"}
	// {"request_id":"65915367-1a4c-4712-b5d1-b2678ef04357","response":{"code":200,"data":"你好,世界!","msg":"请求成功！"},"time":"2022-10-11 09:34:51","url":"/api/v1/demo/demo/index"}
	if json.Valid([]byte(responseData)) {
		var responseJson interface{}
		err = json.Unmarshal([]byte(responseData), &responseJson)
		response = responseJson
	} else {
		response = responseData
	}

	// 记录日志
	log.GetInstance().Log(ctx, logrus.Fields{
		// 请求url(不含域名)
		"url": ctx.Request.URL.Path,
		// 请求参数
		"params": inputs,
		// 返回参数
		"response": response,
	})
}

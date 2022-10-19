// 工具包
// 一些比较公用的工具方法可以集中写在这里,方便使用
// 创建人： 黄翠刚
// 创建时间： 2022.10.09
// 使用示例:utility.Common().FmtPrint()

package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

// ICommon 声明接口类型
type ICommon interface {
	// FmtPrint 方便打印调试
	FmtPrint(data interface{})
	// FileIsExisted 文件或文件夹是否存在
	FileIsExisted(filename string) bool
	// MkDir 新建文件夹
	MkDir(src string) error
	// IsNotExistMkDir 检查文件夹是否存在，如果不存在则新建文件夹
	IsNotExistMkDir(src string) error
	// RequestInputs 获取所有请求参数
	RequestInputs(c *gin.Context) (map[string]interface{}, error)
	// ShouldBind 绑定参数(官方的c.ShouldBind使用一次后,再次获取c.Request.Body的数据时会报错EOF,所以这里取了数据后重新写回c.Request.Body)
	ShouldBind(c *gin.Context, rule any)
}

// 声明结构体类型
type commonImpl struct{}

// Common 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func Common() ICommon {
	return &commonImpl{}
}

// FmtPrint 方便打印调试
func (s *commonImpl) FmtPrint(data interface{}) {
	fmt.Printf("打印开始----------------------------------------------------------------------- \n")
	fmt.Printf("%+v \n", data)
	fmt.Printf("打印结束----------------------------------------------------------------------- \n\n\n")
}

// FileIsExisted 文件或文件夹是否存在
func (s *commonImpl) FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// MkDir 新建文件夹
func (s *commonImpl) MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// IsNotExistMkDir 检查文件夹是否存在，如果不存在则新建文件夹
func (s *commonImpl) IsNotExistMkDir(src string) error {
	if exist := s.FileIsExisted(src); exist == false {
		if err := s.MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// RequestInputs 获取所有请求参数
func (s *commonImpl) RequestInputs(c *gin.Context) (map[string]interface{}, error) {

	const defaultMemory = 32 << 20
	contentType := c.ContentType()

	var (
		dataMap  = make(map[string]interface{})
		queryMap = make(map[string]interface{})
		postMap  = make(map[string]interface{})
	)

	// @see gin@v1.7.7/binding/query.go ==> func (queryBinding) Bind(req *http.Request, obj interface{})
	for k := range c.Request.URL.Query() {
		queryMap[k] = c.Query(k)
	}

	if "application/json" == contentType {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// @see gin@v1.7.7/binding/json.go ==> func (jsonBinding) Bind(req *http.Request, obj interface{})
		if c.Request != nil && c.Request.Body != nil {
			if err := json.NewDecoder(c.Request.Body).Decode(&postMap); err != nil {
				return nil, err
			}
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	} else if "multipart/form-data" == contentType {
		// @see gin@v1.7.7/binding/form.go ==> func (formMultipartBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			return nil, err
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	} else {
		// ParseForm 解析 URL 中的查询字符串，并将解析结果更新到 r.Form 字段
		// 对于 POST 或 PUT 请求，ParseForm 还会将 body 当作表单解析，
		// 并将结果既更新到 r.PostForm 也更新到 r.Form。解析结果中，
		// POST 或 PUT 请求主体要优先于 URL 查询字符串（同名变量，主体的值在查询字符串的值前面）
		// @see gin@v1.7.7/binding/form.go ==> func (formBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseForm(); err != nil {
			return nil, err
		}
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			if err != http.ErrNotMultipart {
				return nil, err
			}
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	}

	var mu sync.RWMutex
	for k, v := range queryMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}
	for k, v := range postMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}

	return dataMap, nil
}

// ShouldBind 绑定参数(官方的c.ShouldBind使用一次后,再次获取c.Request.Body的数据时会报错EOF,所以这里取了数据后重新写回c.Request.Body)
func (s *commonImpl) ShouldBind(c *gin.Context, rule any) {
	bodyBytes, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	_ = c.ShouldBind(&rule)

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

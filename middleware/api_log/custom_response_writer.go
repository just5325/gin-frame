package api_log

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

// 定义一个新的customResponseWriter，通过组合方式持有一个gin.ResponseWriter和response body缓存
type customResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w customResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w customResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

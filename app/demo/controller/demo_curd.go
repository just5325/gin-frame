package controller

import (
	"gin-frame/utility/response"
	"github.com/gin-gonic/gin"
)

var DemoCurd = demoCurdController{}

type demoCurdController struct{}

// Create 数据库create操作
func (c *demoCurdController) Create(ctx *gin.Context) {
	response.Response(ctx).SusJson(nil)
	return
}

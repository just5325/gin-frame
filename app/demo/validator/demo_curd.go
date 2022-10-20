package validator

import (
	utilityValidator "gin-frame/utility/validator"
	"github.com/gin-gonic/gin"
)

var DemoCurd = demoCurdValidator{}

type demoCurdValidator struct{}

// DemoCurdValidatorCreate 为路由 /api/v1/demo/demo_curd/create 设置的接口参数验证规则结构体
type DemoCurdValidatorCreate struct {
	Username string `form:"username"  validate:"required" label:"用户名" validateMsg:"这是自定义验证器错误提示,用户名不能为空"`
	Password string `json:"password" validate:"required" label:"密码"`
}

// Create 为路由 /api/v1/demo/demo_curd/create 设置的接口参数验证方法
func (c *demoCurdValidator) Create(ctx *gin.Context) {
	var data DemoCurdValidatorCreate
	utilityValidator.Validator(ctx).Struct(&data)
	return
}

// DemoCurdValidatorUpdate 为路由 /api/v1/demo/demo_curd/update 设置的接口参数验证规则结构体
type DemoCurdValidatorUpdate struct {
	Username string `form:"username"  validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required,demo" label:"密码"`
}

// Update 为路由 /api/v1/demo/demo_curd/update 设置的接口参数验证方法
func (c *demoCurdValidator) Update(ctx *gin.Context) {
	var data DemoCurdValidatorUpdate
	utilityValidator.Validator(ctx).Struct(&data)
	return
}

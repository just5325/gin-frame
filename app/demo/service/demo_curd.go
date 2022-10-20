package service

import (
	"gin-frame/app/demo/dao/model"
	"gin-frame/app/demo/dao/query"
	"gin-frame/dao/db"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"time"
)

// IDemoCurd 声明接口类型
type IDemoCurd interface {
	// Create 数据库create操作
	Create(username string, password string) (res gin.H, err error)
	// Update 数据库update操作
	Update(username string, password string) (err error)
	// CheckUsernameExist 判断指定用户名是否存在
	CheckUsernameExist(username string) (res bool, err error)
}

// 声明结构体类型
type demoCurdImpl struct {
	ctx *gin.Context
}

// DemoCurd 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func DemoCurd(ctx *gin.Context) IDemoCurd {
	return &demoCurdImpl{
		ctx: ctx,
	}
}

// Create 数据库create操作
func (s *demoCurdImpl) Create(username string, password string) (res gin.H, err error) {
	defaultDb, err := db.Db().GetDb("default")
	if err != nil {
		return
	}
	modelUser := model.User{
		Username:   username,
		Password:   password,
		State:      1,
		CreateTime: cast.ToInt32(time.Now().Unix()),
	}

	// 创建 +++++++++++++++++++++
	// 创建记录
	// 方式1: 使用 gorm包
	//getDb.Create(&modelUser)

	// 方式2: 使用 gorm.io/gen包
	u := query.Use(defaultDb).User
	err = u.Debug().Create(&modelUser)
	if err != nil {
		return
	}

	// 方式3: 选择字段创建
	// 创建记录并为指定的字段赋值。
	//u := query.Use(defaultDb).User
	//err = u.WithContext(s.ctx).Select(u.Username, u.Password).Create(&modelUser)
	//if err != nil {
	//	return
	//}
	// INSERT INTO `user` (`username`,`password`) VALUES ("1", "1")

	res = gin.H{
		"id": cast.ToInt(modelUser.ID),
	}
	return
}

// Update 数据库update操作
func (s *demoCurdImpl) Update(username string, password string) (err error) {
	// CheckUsernameExist 判断指定用户名是否存在
	checkUsernameExist, err := s.CheckUsernameExist(username)
	if err != nil {
		return
	}
	if !checkUsernameExist {
		err = errors.New("用户名不存在")
		return
	}

	defaultDb, err := db.Db().GetDb("default")
	if err != nil {
		return
	}
	u := query.Use(defaultDb).User

	//// 更新单个字段
	//_, err = u.Where(u.Username.Eq(username)).Update(u.Password, password)

	// 更新多个字段
	//_, err = u.Where(u.Username.Eq(username)).Updates(map[string]interface{}{"password": password, "state": 1})
	_, err = u.Where(u.Username.Eq(username)).Updates(model.User{Password: password, State: 1})

	return
}

// CheckUsernameExist 判断指定用户名是否存在
func (s *demoCurdImpl) CheckUsernameExist(username string) (res bool, err error) {
	defaultDb, err := db.Db().GetDb("default")
	if err != nil {
		return
	}

	u := query.Use(defaultDb).User
	count, err := u.Where(u.Username.Eq(username)).Count()
	if err != nil {
		return
	}

	if count >= 1 {
		res = true
	}

	return
}

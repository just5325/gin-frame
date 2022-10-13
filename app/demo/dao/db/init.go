// 数据库连接实例
// 设计模式：简单工厂 + 单例模式
// 使用示例: db.Db().GetDb("cd")
// 创建人： 黄翠刚
// 创建时间： 2022.10.13

package db

import (
	"fmt"
	"gin-frame/config"
	"gin-frame/utility/log"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var dbMap = make(map[string]*gorm.DB)

type dbConfig struct {
	host        string
	port        string
	user        string
	passwd      string
	dbname      string
	maxOpenConn int
	maxIdleConn int
}

// IDb 声明接口类型
type IDb interface {
	GetDb(dbName string) (*gorm.DB, error)
}

// 声明结构体类型
type dbImpl struct{}

// Db 声明一个方法，用于获取当前包主要结构体的对象，便于执行其方法
func Db() IDb {
	return &dbImpl{}
}

// GetDb 方法
func (s *dbImpl) GetDb(dbName string) (*gorm.DB, error) {
	if value, ok := dbMap[dbName]; ok {
		return value, nil
	}

	// 获取数据库的配置信息
	configPath := "database." + dbName
	database := config.Config().GetViper().GetStringMap(configPath)
	if len(database) == 0 {
		return nil, errors.New("缺少数据库配置信息:" + configPath)
	}

	configData := dbConfig{
		host:        cast.ToString(database["host"]),
		port:        cast.ToString(database["port"]),
		user:        cast.ToString(database["user"]),
		passwd:      cast.ToString(database["passwd"]),
		dbname:      cast.ToString(database["dbname"]),
		maxOpenConn: cast.ToInt(database["max_open_cons"]),
		maxIdleConn: cast.ToInt(database["max_idle_cons"]),
	}

	var once sync.Once
	var initDBErr error
	once.Do(func() {
		dbMap[dbName], initDBErr = initDB(configData)
	})

	return dbMap[dbName], initDBErr
}

func initDB(config dbConfig) (*gorm.DB, error) {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.user, config.passwd, config.host, config.port, config.dbname)

	myLogger := logger.New(
		log.NewMyWriter(),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢SQL阈值
			LogLevel:                  logger.Info,            // 设置日志级别，只有Info以上才会打印sql
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                  // 关闭彩打
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: myLogger.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxIdleConns(config.maxOpenConn)
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(config.maxIdleConn)

	return db.Debug(), nil
}

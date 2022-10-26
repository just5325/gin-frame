# gin-frame
使用gin搭建一个快速开发的框架,仅满足我自己使用即可

### 项目目录结构规划
```
/ 项目根目录
├── app 应用模块
│   └── demo1 demo模块1
│       ├── controller 控制器目录
│       ├── dao 数据访问层（Data Access Layer），用于和底层数据库交互，仅包含最基础的 CURD 方法
│       │   ├── query （gorm.io/gen）自动生成 CRUD 和 DIY 方法的代码的目录
│       │   └── model （gorm.io/gen）自动生成 model 文件的目录
│       ├── router 路由目录
│       ├── validator 路由验证器目录
│       └── service 服务层目录
├── cmd 可执行文件目录
├── config 配置文件目录
│   └── config.yml 项目配置文件
├── router 路由
├── middleware 中间件目录
├── dao 数据访问层（Data Access Layer），用于和底层数据库交互，仅包含最基础的 CURD 方法
│   ├── db 数据库连接,用于创建与数据库的连接
│   ├── query （gorm.io/gen）自动生成 CRUD 和 DIY 方法的代码的目录
│   └── model （gorm.io/gen）自动生成 model 文件的目录
├── resource 资源目录
│   ├── image 图片
│   └── logs 日志文件
├── utility 工具包
│   ├── log 日志工具包
│   ├── response 返回数据工具包
│   ├── validator 验证器工具包
│   └── common.go 公共方法,一些比较公用的工具方法可以集中写在这里,方便使用
├── go.mod 使用 Go module 包管理的依赖描述文件
└── main.go 程序入口文件
```

### 项目依赖的三方库
```
基础框架: github.com/gin-gonic/gin
数据库: gorm.io/gorm + gorm.io/driver/mysql
数据库模型自动生成: gorm.io/gen
配置文件: github.com/spf13/viper
类型转换: github.com/spf13/cast
日志: github.com/sirupsen/logrus
日志分割: github.com/lestrrat-go/file-rotatelogs + github.com/rifflock/lfshook
全局请求ID中间件: github.com/gin-contrib/requestid
Redis缓存: github.com/go-redis/redis/v9
```

### 项目部署目录结构规划
```
/ 项目根目录
├── config 配置文件目录
│   └── config.yml 项目配置文件
├── resource 资源目录
│   └── logs 日志文件(可通过config.yml更改日志文件路径等等配置)
├── pid PID文件(记录了程序运行进程的pid,用于执行kill pid使用,注意:不要使用kill -9 pid,因为程序中捕获了程序中断的系统信号并且做了优雅关机的处理)
└── mian 可执行二进制文件(注意!注意!注意!上述所有目录及文件都可由程序运行时自动创建,无需手动创建目录结构,还有注意配合守护进程使用)
```


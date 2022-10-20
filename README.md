# gin-frame
使用gin搭建一个快速开发的框架,仅满足我自己使用即可

### 项目目录结构规划
```
/
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
```

### 代码设计规范
```

```


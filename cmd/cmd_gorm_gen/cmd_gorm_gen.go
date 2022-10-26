package cmd_gorm_gen

import (
	"fmt"
	"gin-frame/config"
	"gin-frame/dao/db"
	"github.com/spf13/cast"
	"gorm.io/gen"
)

type gormGenConfigT struct {
	// isSetConfig 是否有配置文件配置当前配置项
	isSetConfig bool
	// out_path 指定查询代码路径 (默认:"./query")
	outPath string
	// model_pkg_path 指定模型代码路径 (默认:"./model")
	modelPkgPath string
	// with_unit_test 为查询代码生成单元测试 (默认:false)
	withUnitTest bool
}

type genModelT struct {
	// database 数据库（如：”default“，表示使用本配置文件中 database.default 数据库）
	database string
	// table 表名称
	table         []string
	gormGenConfig gormGenConfigT
}

type gormGenT struct {
	// start 开关 (true:随着主协程启动时完成生成模型的操作 (默认)false:不执行)
	start         bool
	gormGenConfig gormGenConfigT
	genModel      []genModelT
}

var gormGen gormGenT

func init() {
	// 初始化配置信息
	initConfig()
	// 根据判断开关配置判断是否执行生成模型文件的操作(true:完成生成模型的操作 false:不执行)
	if !gormGen.start {
		return
	}

	// 遍历要生成模型文件的配置数组
	for _, value := range gormGen.genModel {
		// 配置项
		newGeneratorConfig := gen.Config{
			// 如果你想查询没有上下文约束，设置模式 gen.WithoutContext
			Mode: gen.WithoutContext,
		}

		// 如果遍历的当前配置项有设置过 gormGenConfig,则使用当前配置项的设置信息作为配置
		// 如果遍历的当前配置项没有设置过 gormGenConfig,则使用主配置的 gormGenConfig 配置(如果都没有配置过gormGenConfig也会有默认配置的)
		if value.gormGenConfig.isSetConfig {
			newGeneratorConfig.OutPath = value.gormGenConfig.outPath
			newGeneratorConfig.ModelPkgPath = value.gormGenConfig.modelPkgPath
			newGeneratorConfig.WithUnitTest = value.gormGenConfig.withUnitTest
		} else {
			newGeneratorConfig.OutPath = gormGen.gormGenConfig.outPath
			newGeneratorConfig.ModelPkgPath = gormGen.gormGenConfig.modelPkgPath
			newGeneratorConfig.WithUnitTest = gormGen.gormGenConfig.withUnitTest
		}

		execute(value.database, value.table, newGeneratorConfig)
	}
}

func getGormGenConfig(configPath string) gormGenConfigT {
	gormGenConfig := gormGenConfigT{
		// out_path 指定查询代码路径 (默认:"./query")
		outPath: config.GetInstance().GetViper().GetString(fmt.Sprintf("%s.out_path", configPath)),
		// mode l_pkg_path 指定模型代码路径 (默认:"./model")
		modelPkgPath: config.GetInstance().GetViper().GetString(fmt.Sprintf("%s.model_pkg_path", configPath)),
		// with_unit_test 为查询代码生成单元测试 (默认:false)
		withUnitTest: config.GetInstance().GetViper().GetBool(fmt.Sprintf("%s.with_unit_test", configPath)),
	}
	// 判断是否有配置文件配置当前配置项(只要有一项不为空，则说明有配置文件配置当前配置项)
	gormGenConfig.isSetConfig = gormGenConfig.outPath != "" || gormGenConfig.modelPkgPath != "" || gormGenConfig.withUnitTest != false

	if gormGenConfig.outPath == "" {
		gormGenConfig.outPath = "dao/query"
	}
	if gormGenConfig.modelPkgPath == "" {
		gormGenConfig.modelPkgPath = "dao/model"
	}

	return gormGenConfig
}

// 初始化配置信息
func initConfig() {
	genModel := make([]genModelT, 0)
	//var genModel = []genModelT
	genModelToSlice := cast.ToSlice(config.GetInstance().GetViper().Get("gorm_gen.gen_model"))
	for v := range genModelToSlice {
		configPath := fmt.Sprintf("gorm_gen.gen_model.%d", v)
		genModelItem := genModelT{
			database:      config.GetInstance().GetViper().GetString(fmt.Sprintf("%s.database", configPath)),
			table:         config.GetInstance().GetViper().GetStringSlice(fmt.Sprintf("%s.table", configPath)),
			gormGenConfig: getGormGenConfig(fmt.Sprintf("%s.gorm_gen_config", configPath)),
		}

		genModel = append(genModel, genModelItem)
	}

	// 获取生成模型的配置数组
	gormGen = gormGenT{
		start:         config.GetInstance().GetViper().GetBool("gorm_gen.start"),
		gormGenConfig: getGormGenConfig("gorm_gen.gorm_gen_config"),
		genModel:      genModel,
	}
}

// 执行代码生成的操作
func execute(database string, table []string, newGeneratorConfig gen.Config) {
	// 在项目中重用数据库连接或在这里创建一个连接
	// 如果你想使用 GenerateModel/GenerateModelAs , UseDB是必要的，否则它会 panic
	gormDb, err := db.GetInstance().GetDb(database)
	if err != nil {
		panic(err)
	}

	generator := gen.NewGenerator(newGeneratorConfig)

	generator.UseDB(gormDb)

	// 在结构或表模型上应用基本的crud API，由表名和函数指定
	// GenerateModel / GenerateModelAs。生成器将在调用execute时生成表模型的代码。
	var models []interface{}
	for v := range table {
		models = append(models, generator.GenerateModel(table[v]))
	}

	generator.ApplyBasic(models...)

	// 执行代码生成的操作
	generator.Execute()
}

package bootstrap

import (
	"fmt"
	"github.com/crack007/api-base/common/config"
	"github.com/crack007/api-base/core"
	"github.com/crack007/api-base/route"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func loadDefaultConfig(app *core.App) {
	// TODO 默认配置
}
func Init(app *core.App) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("获取RootPath失败：", err)
	}
	app.RootPath = dir
	InitConfig(app)
	InitDB(app)
	InitRouter(app)
	if app.InitResource != nil {
		app.InitResource(app)
	}
}
func parseCommandArgs() {
	pflag.String("configName", "", "配置文件名")
	pflag.String("configType", "", "配置文件类型(json,toml,yaml,hcl,env和java properties 配置类型)")
	pflag.String("configPath", "", "配置文件路径")
	pflag.Parse()
}
func InitConfig(app *core.App) {
	parseCommandArgs()
	bindPFlagsErr := viper.BindPFlags(pflag.CommandLine)
	if bindPFlagsErr != nil {
		log.Println("BindPFlags Err:", bindPFlagsErr)
	}
	gin.DisableConsoleColor()
	var configName = viper.GetString("configName")
	if configName == "" {
		configName = "appConfig"
	}
	var configType = viper.GetString("configType")
	if configType == "" {
		configType = "yml"
	}
	var configPath = viper.GetString("configPath")
	if configPath == "" {
		configPath = app.RootPath + "/config"
	}
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		loadDefaultConfig(app)
	}
	InitDBConfig(app)
	InitCommonConfig(app)
	if app.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}
	app.Engine = gin.Default()
}

func InitCommonConfig(app *core.App) {
	if app.InitConfig != nil {
		app.InitConfig(app)
	}
	config.GetCommonConfig().SetEnv(viper.GetString("app.env"))
	config.GetCommonConfig().SetHttpTimeout(viper.GetInt("app.httpTimeout"))
	if app.IsDev() {
		viper.Debug()
	}
	log.Println("InitCommonConfig... ok!")
}
func InitDBConfig(app *core.App) {
	dbConfig := config.GetDbConfig()
	dbConfig.SetEngine(viper.GetString("db.engine"))
	dbConfig.SetHost(viper.GetString("db.host"))
	dbConfig.SetUser(viper.GetString("db.user"))
	dbConfig.SetDatabase(viper.GetString("db.database"))
	dbConfig.SetPort(viper.GetInt("db.port"))
	dbConfig.SetPassword(viper.GetString("db.password"))
	dbConfig.SetCharset(viper.GetString("db.charset"))
}
func InitDB(app *core.App) {
	dbConfig := config.GetDbConfig()
	var url = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User(),
		dbConfig.Password(),
		dbConfig.Host(),
		dbConfig.Port(),
		dbConfig.Database(),
		dbConfig.Charset(),
	)
	db, err := gorm.Open(dbConfig.Engine(), url)
	if err != nil {
		log.Println("db url=", url)
		log.Fatal("db connect error", err)
	}
	db.DB().SetMaxIdleConns(dbConfig.MaxIdleConnections())
	db.DB().SetMaxOpenConns(dbConfig.MaxOpenConnections())
	app.Db = db
}
func InitRouter(app *core.App) {
	var r = route.Route{}
	r.Init(app.Engine)
}
func Run(app *core.App) {
	var err = app.Engine.Run(":" + strconv.Itoa(viper.GetInt("app.port")))
	if err != nil {
		log.Fatal("启动失败：", err.Error())
	}
}

# api框架

## 使用说明


- 定义控制器

```go

package controller

import (
	"github.com/crack007/api-base/common/exception"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (index *IndexController) Login(context *gin.Context) interface{} {
	var apiDTO = &dtoApi.LoginApiDTO{}
	var e = context.ShouldBindJSON(apiDTO)
	if e != nil {
		return exception.CreateValidationException(e.Error())
	}
    // TODO 执行业务,返回结果
	return nil
}


```



- 定义路由

```go

package route

import (
	commonMiddleware "github.com/crack007/api-base/common/middleware"
	"github.com/crack007/api-base/route"
	"github.com/gin-gonic/gin"
)

var index = &controller.IndexController{}

func ApiV1() {
	route.CreateGroup("/v1/index", func(routerGroup *gin.RouterGroup) {
    		// 添加中间件
    		routerGroup.Use(commonMiddleware.ResponseMiddleware)
    		// 添加路由,注意使用代理路由组件
    		routerGroup.POST("/login", (&route.AgentRoute{Target: index.Login}).Agent)
    	})
}


```

- 自定义配置

> 配置文件路径默认 运行路径/config/config.yml

> 配置文件路径可以通过启动参数定义，启动参数如下

```

--configName 参数值   配置文件名
--configPath 参数值   配置文件路径
--configType 参数值   配置文件类型(json,toml,yaml,hcl,env和java properties 配置类型)

```

```go

package config

import (
	"github.com/crack007/api-base/common/constant"
	"github.com/spf13/viper"
	"log"
)

var appConfig = &AppConfig{}

func GetAppConfig() *AppConfig {
	return appConfig
}
func init() {
	// 设置默认值
	viper.SetDefault("app.mode", "release")
	viper.SetDefault("app.env", constant.ENV_DEV)
	viper.SetDefault("app.jwtRefreshTokenExpired", 30*24*3600)
	log.Println("初始化默认配置")
}

func InitAppConfig() {
	// TODO 设置自定义配置
	GetAppConfig().SetJwtRefreshTokenExpired(viper.GetInt("app.jwtRefreshTokenExpired"))
}

type AppConfig struct {
	jwtExpired             int    `desc:"jwt有效期(单位：秒)"`
	jwtRefreshTokenExpired int    `desc:"jwt刷新token有效期(单位：秒)"`
}

func (a *AppConfig) JwtRefreshTokenExpired() int {
	return a.jwtRefreshTokenExpired
}

func (a *AppConfig) SetJwtRefreshTokenExpired(jwtRefreshTokenExpired int) {
	a.jwtRefreshTokenExpired = jwtRefreshTokenExpired
}


```

- 入口文件引入

```go

import (
	bootstrap "github.com/crack007/api-base"
	"github.com/crack007/api-base/core"
)
var app = core.GetApp()
app.InitConfig = func(app *core.App) {
    // TODO 初始化配置
   config.InitAppConfig()
}
app.InitRouter = func(app *core.App) {
    // TODO 引用路由
   route.ApiV1()
}
app.InitResource = func(app *core.App) {
    // TODO 初始化资源对象等
}
bootstrap.Init(app)
bootstrap.Run(app)

```

- 引用gorm实例

```go

core.GetApp().Db

```
# 默认配置参考

```yaml

db:
  engine: mysql
# ip地址
  host:
# 端口
  port:
# 账号
  user:
# 密码
  password:
# 数据库名
  database:
# 链接字符集
  charset:
# 最大空闲连接数
  maxIdleConnections:
# 最大打开连接数
  maxOpenConnections:
app:
  # http监听端口
  port: 80
  # 运行模式
  mode: debug

```

# TODO-LIST

- 支持配置路径自定义 [已完成]

- 缓存支持 [待完成]
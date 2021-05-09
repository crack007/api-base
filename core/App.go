package core

import (
	"github.com/crack007/api-base/common/config"
	"github.com/crack007/api-base/common/constant"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type App struct {
	Engine       *gin.Engine
	Db           *gorm.DB
	RootPath     string
	InitConfig   callback
	InitRouter   callback
	InitResource callback
}
type callback func(app *App)

var app = &App{}

func GetApp() *App {
	return app
}

func (a *App) IsProd() bool {
	return config.GetCommonConfig().Env() == constant.ENV_PROD
}

func (a *App) IsTest() bool {
	return config.GetCommonConfig().Env() == constant.ENV_TEST
}

func (a App) IsDev() bool {
	return config.GetCommonConfig().Env() == constant.ENV_DEV
}

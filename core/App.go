package core

import (
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

package route

import (
	"cracker.com/base/core"
	"github.com/gin-gonic/gin"
)

type Route struct {
}

func (r *Route) Init(e *gin.Engine) {
	// 添加路由
	if core.GetApp().InitRouter != nil {
		core.GetApp().InitRouter(core.GetApp())
	}
}

func CreateGroup(path string, _callback callback) {
	_callback(core.GetApp().Engine.Group(path))
}

type callback func(routerGroup *gin.RouterGroup)

type InitRoute interface {
	init(e *gin.Engine)
}

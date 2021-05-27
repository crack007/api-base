package route

import (
	"github.com/crack007/api-base/core"
	"github.com/gin-gonic/gin"
)

type Route struct {
}
type Group struct {
	RouterGroup *gin.RouterGroup
}

func (r *Route) Init(e *gin.Engine) {
	// 添加路由
	if core.GetApp().InitRouter != nil {
		core.GetApp().InitRouter(core.GetApp())
	}
}

func CreateGroup(path string) *Group {
	var g = &Group{
		RouterGroup: core.GetApp().Engine.Group(path),
	}
	return g
}

type InitRoute interface {
	init(e *gin.Engine)
}

// 创建路由代理目标对象
func createAgent(target func(ctx *gin.Context) interface{}) func(context *gin.Context) {
	return (&AgentRoute{Target: target}).Agent
}

func (r *Group) CreateRoute(method string, url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.Handle(method, url, createAgent(action))
	return r
}

func (r *Group) GET(url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.GET(url, createAgent(action))
	return r
}

func (r *Group) POST(url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.POST(url, createAgent(action))
	return r
}

func (r *Group) PUT(url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.PUT(url, createAgent(action))
	return r
}

func (r *Group) PATCH(url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.PATCH(url, createAgent(action))
	return r
}

func (r *Group) DELETE(url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.DELETE(url, createAgent(action))
	return r
}

func (r *Group) OPTIONS(url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.OPTIONS(url, createAgent(action))
	return r
}

func (r *Group) HEAD(url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.HEAD(url, createAgent(action))
	return r
}

func (r *Group) Any(url string, action func(ctx *gin.Context) interface{}) *Group {
	r.RouterGroup.Any(url, createAgent(action))
	return r
}

func (r *Group) Use(middleware gin.HandlerFunc) *Group {
	r.RouterGroup.Use(middleware)
	return r
}

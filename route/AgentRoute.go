package route

import "github.com/gin-gonic/gin"

type AgentRoute struct {
	Target func(ctx *gin.Context) interface{}
}

func (agent *AgentRoute) Agent(context *gin.Context) {
	context.Keys["_response"] = agent.Target(context)
}

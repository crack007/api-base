package middleware

import (
	"github.com/crack007/api-base/common/exception"
	"github.com/gin-gonic/gin"
)

func ResponseMiddleware(context *gin.Context) {
	if context.Keys == nil {
		context.Keys = map[string]interface{}{}
	}
	context.Next()
	dto := context.Keys["_response"]
	if v, ok := dto.(*exception.ApiException); ok {
		responseError(context, v.Code(), v.Error())
		return
	}
	if v, ok := dto.(*exception.ValidationException); ok {
		responseError(context, v.Code(), v.Error())
		return
	}
	if v, ok := dto.(*exception.NullException); ok {
		responseError(context, v.Code(), v.Error())
		return
	}
	if v, ok := dto.(*exception.ForbiddenException); ok {
		responseError(context, v.Code(), v.Error())
		return
	}
	context.JSON(200, gin.H{
		"msg":  "",
		"code": 200,
		"data": dto,
	})
}
func responseError(context *gin.Context, code int, err string) {
	context.JSON(200, gin.H{
		"msg":  err,
		"code": code,
		"data": nil,
	})
}

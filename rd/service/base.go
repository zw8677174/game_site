package service

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rd/pkg/app"
	"rd/pkg/e"
)

type Map map[string]interface{}

type BaseService struct {
}

func (t *BaseService) Success(data interface{}) () {
	app.Response(http.StatusOK, e.SUCCESS, data)
	return
}

func (t *BaseService)Param(ctx *gin.Context ) Map {
	var params Map
	err2 := ctx.ShouldBindQuery(&params)
	if err2 ==nil {
	}
	return params
}

func (t *BaseService)Form(ctx *gin.Context ) Map {
	var form Map
	err2 := ctx.ShouldBindJSON(&form)
	if err2 ==nil {
	}
	return form
}


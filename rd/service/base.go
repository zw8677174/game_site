package service

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rd/pkg/app"
	"rd/pkg/e"
)


type BaseService struct {
}

func (t *BaseService) Success(data interface{}) () {
	c := new(gin.Context)
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}


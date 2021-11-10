package service

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rd/pkg/app"
	"rd/pkg/e"
	"strconv"
)

type Map map[string]interface{}

type PageParam struct {
	pageNo int
	pageNum int
}



type BaseService struct {
}

func (t *BaseService) Success(c *gin.Context,data interface{}) () {
	app.Response(c, http.StatusOK, e.SUCCESS, data)
	return
}

func (t *BaseService)Form(ctx *gin.Context ) Map {
	var form Map
	err2 := ctx.ShouldBindJSON(&form)
	if err2 ==nil {
	}
	return form
}

func (t *BaseService)getPageParams(ctx *gin.Context ) PageParam {
	var pageParam PageParam
	pageNoStr, _ := ctx.GetQuery("pageNo")
	pageNumStr , _ := ctx.GetQuery("pageNum")

	pageParam.pageNo, _ = strconv.Atoi(pageNoStr)
	pageParam.pageNum, _ = strconv.Atoi(pageNumStr)
	return pageParam
}




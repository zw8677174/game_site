package app

import (
	"github.com/gin-gonic/gin"
	"rd/pkg/e"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *gin.Context) Response(httpCode, errCode int, data interface{}) {
	g.JSON(httpCode, ResponseData{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rd/service"

	"rd/middleware/jwt"
	"rd/pkg/export"
	"rd/pkg/qrcode"
	"rd/pkg/upload"
	"rd/routers/api"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/api/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 游戏列表
		apiv1.GET("/game_list", new(service.GameService).GetList)

	}

	return r
}

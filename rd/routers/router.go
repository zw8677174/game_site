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

	apiFont := r.Group("/api/font")
	apiFont.Use(jwt.JWT())
	{
		apiFont.GET("/game_list", new(service.GameService).GetList)
	}

	apiBackend := r.Group("/api/bd")
	apiBackend.Use(jwt.JWT())
	{
		apiBackend.GET("/game_list", new(service.GameService).GetAuthorList)
		apiBackend.GET("/game_create", new(service.GameService).Create)

	}

	return r
}

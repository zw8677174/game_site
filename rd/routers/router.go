package routers

import (
	"github.com/gin-gonic/gin"
	"rd/service"

	"rd/middleware/jwt"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/api/auth", new(service.AuthService).Auth)

	apiFont := r.Group("/api/front")
	apiFont.Use(jwt.JWT())
	{
		apiFont.GET("/game_list", new(service.GameService).GetList)
	}

	apiBackend := r.Group("/api/bd")
	apiBackend.Use(jwt.JWT())
	{
		apiBackend.GET("/game_list", new(service.GameService).GetAuthorList)
		apiBackend.POST("/game_create", new(service.GameService).Create)
	}

	return r
}

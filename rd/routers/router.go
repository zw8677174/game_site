package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rd/middleware/jwt"
	"rd/pkg/export"
	"rd/pkg/qrcode"
	"rd/pkg/upload"
	"rd/routers/api"
	"rd/routers/api/v1"
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
	r.GET("/tags1", v1.GetTags)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 游戏列表
		apiv1.GET("/game_list", v1.GetGameList)
        //获取标签列表
		apiv1.GET("/tags", v1.GetTags)

		//新建标签
        apiv1.POST("/tags", v1.AddTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}

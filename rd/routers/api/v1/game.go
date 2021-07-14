package v1

import (
	"rd/service"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"rd/pkg/app"
	"rd/pkg/setting"
	"rd/pkg/util"
)


func GetGameList(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	gameService := service.GameService{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

}

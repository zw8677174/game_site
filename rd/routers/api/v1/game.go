package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"rd/pkg/app"
	"rd/pkg/e"
	"rd/pkg/setting"
	"rd/pkg/util"
	"rd/service/game_service"
)


func GetGameList(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	gameService := game_service.Game{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := gameService.GetList()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tags,
	})
}

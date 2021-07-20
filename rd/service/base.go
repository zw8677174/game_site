package service

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rd/models"
	"rd/pkg/app"
	"rd/pkg/e"
)


type GameService struct {
	BaseService
}

func (t *GameService) GetList(c *gin.Context) () {
	appG := app.Gin{C: c}

	var (
		Games []models.Game
	)

	Games, err := new(models.Game).GetList(t.PageInfo.PageNum, t.PageInfo.PageSize, t.getCommonConds())
	if err != nil {
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": Games,
	})

}

func (t *GameService) GetAuthorList(c *gin.Context) () {
	appG := app.Gin{C: c}

	conds := t.getCommonConds()
	conds["author_id"] = c.GetInt64("uid")
	Games, err := new(models.Game).GetList(t.PageInfo.PageNum, t.PageInfo.PageSize, conds)

	if err != nil {
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": Games,
	})

}

func (t *GameService) Create(c *gin.Context) () {
	inputMap := c.GetStringMap("json")
	appG := app.Gin{C: c}
	game := models.Game{
		Name: inputMap["name"].(string),
		AuthorId: c.GetInt64("uid"),
	}
	err := game.Create

	if err != nil {
	}
	
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"info": game,
	})

}

func (t *GameService) getCommonConds() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_del"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}


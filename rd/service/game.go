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

	Games, err := new(models.Game).GetList(c.GetStringMap("pageInfo")["pageNo"].(int), c.GetStringMap("pageInfo")["pageInfo"].(int), t.getCommonConds())
	if err != nil {
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": Games,
	})

}

func (t *GameService) GetAuthorList(c *gin.Context) () {

	conds := t.getCommonConds()
	conds["author_id"] = c.GetInt64("uid")
	Games, err := new(models.Game).GetList(c.GetStringMap("pageInfo")["pageNo"].(int), c.GetStringMap("pageInfo")["pageInfo"].(int), conds)

	if err != nil {
	}

	t.Success(Games)

}

func (t *GameService) Create(c *gin.Context) () {
	inputMap := c.GetStringMap("json")
	game := models.Game{
		Name:     inputMap["name"].(string),
		AuthorId: c.GetInt64("uid"),
	}
	err := game.Create

	if err != nil {
	}
	t.Success(nil)

}

func (t *GameService) getCommonConds() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_del"] = 0
	return maps
}

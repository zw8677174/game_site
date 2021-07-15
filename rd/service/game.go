package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rd/models"
	"rd/pkg/app"
	"rd/pkg/e"
)

type GameService struct {
	ID         int
	Name       string
	CreatedBy  string
	State  int

	PageNum  int
	PageSize int
}

func (t *GameService) GetList(c *gin.Context) () {
	appG := app.Gin{C: c}

	var (
		Games []models.Game
	)

	Games, err := new(models.Game).GetList(t.PageNum, t.PageSize, t.getListConds())
	if err != nil {
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": Games,
	})

}

func (t *GameService) getListConds() map[string]interface{} {
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


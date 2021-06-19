package game_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
)

type Game struct {
	ID         int
	Name       string
	CreatedBy  string
	State  int

	PageNum  int
	PageSize int
}

func (t *Game) GetList() ([]models.Game, error) {
	var (
		tags []models.Game
	)

	tags, err := models.GetList(t.PageNum, t.PageSize, t.getConds())
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *Game) getConds() map[string]interface{} {
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


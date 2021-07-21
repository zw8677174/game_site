package service

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rd/models"
)

type GameService struct {
	BaseService
}

func (t *GameService) GetList(c *gin.Context) () {
	Games, err := new(models.Game).GetList(c.GetStringMap("pageInfo")["pageNo"].(int), c.GetStringMap("pageInfo")["pageInfo"].(int), t.getCommonConds())
	if err != nil {
	}
	t.Success(Games)
}

func (s *GameService) GetAuthorList(c *gin.Context) () {

	conds := s.getCommonConds()
	conds["author_id"] = c.GetInt64("uid")
	Games, _ := new(models.Game).GetList(c.GetStringMap("pageInfo")["pageNo"].(int), c.GetStringMap("pageInfo")["pageInfo"].(int), conds)
	s.Success(Games)
}

func (t *GameService) Create(c *gin.Context) () {
	input := t.Form(c)

	fmt.Print(input)
	game := models.Game{
		Name: input["name"].(string),
		AuthorId: c.GetInt64("uid"),
	}
	game.Create()
	t.Success(nil)
	return
}

func (t *GameService) getCommonConds() map[string]interface{} {
	conds := make(map[string]interface{})
	conds["is_del"] = 0
	return conds
}

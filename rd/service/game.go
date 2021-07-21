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

func (t *GameService) GetAuthorList(c *gin.Context) () {

	conds := t.getCommonConds()
	conds["author_id"] = c.GetInt64("uid")
	Games, _ := new(models.Game).GetList(c.GetStringMap("pageInfo")["pageNo"].(int), c.GetStringMap("pageInfo")["pageInfo"].(int), conds)
	t.Success(Games)
}

func (t *GameService) Create(c *gin.Context) () {
	inputMap := c.GetStringMap("json")
	fmt.Print(inputMap)
	fmt.Print("*****")

	//game := models.Game{
	//	Name:     inputMap["name"].(string),
	//	AuthorId: c.GetInt64("uid"),
	//}
	//game.Create()
	//t.Success(nil)
}

func (t *GameService) getCommonConds() map[string]interface{} {
	conds := make(map[string]interface{})
	conds["is_del"] = 0
	return conds
}

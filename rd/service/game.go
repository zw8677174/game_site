package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rd/models"
	"rd/pkg/app"
)

type GameService struct {
	BaseService
}

func (t *GameService) GetList(c *gin.Context) () {
	pageParams := t.getPageParams(c)
	conds := t.getCommonConds()
	name := t.getQueryString(c, "name")
	if name != "" {
		conds["name"] = name
	}
	Games := new(models.Game).GetList(pageParams.pageNo, pageParams.pageNum, conds)

	t.Success(c, Games)
}

func (s *GameService) GetAuthorList(c *gin.Context) () {
	conds := s.getCommonConds()
	conds["author_id"] = c.GetInt64("uid")
	//Games, _ := new(models.Game).GetList(c.GetStringMap("pageInfo")["pageNo"].(int), c.GetStringMap("pageInfo")["pageInfo"].(int), conds)
	//s.Success(c, Games)
}

func (gameServer *GameService) Create(c *gin.Context) () {
	input := gameServer.getFormMap(c)

	fmt.Print(input)
	game := models.Game{
		Name: input["name"].(string),
		AuthorId: c.GetInt64("uid"),
	}
	game.Create()

	newC := new(gin.Context)
	newC.JSON(200,app.ResponseData{
		Code: 200,
		Data: nil,
	})
	gameServer.Success(c,nil)
	return
}

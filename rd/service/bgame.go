package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rd/models"
	"rd/pkg/app"
	"rd/pkg/e"
)

type BGameService struct {
	ID        int
	Name      string
	CreatedBy string
	State     int

	PageNum  int
	PageSize int
}

func (t *BGameService) Create(c *gin.Context) () {
	appG := app.Gin{C: c}

	err := new(models.Game).create("test")
	if err != nil {
	}

	appG.Response(http.StatusOK, e.SUCCESS, true)
}

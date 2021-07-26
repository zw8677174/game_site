package service

import (
	"github.com/gin-gonic/gin"
	"rd/models"
)

type AuthService struct {
	Username string
	Password string
}

func (a *AuthService) GetUid(c *gin.Context)  int64 {
	user := new(models.User).Get(c.GetStringMap("pageInfo")["pageNo"].(int))
	return user.Uid
}
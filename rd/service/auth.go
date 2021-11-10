package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rd/models"
)

type AuthService struct {
	Username string
	Password string
}

func (a *AuthService) GetUid(c *gin.Context)  int64 {
	conds := make(map[string]interface{})
	conds["username"] = a.Username
	conds["password"] = a.Password
	user := new(models.User).GetOne(conds)
	fmt.Print(user)
	return int64(1)
	//return user.Uid
}
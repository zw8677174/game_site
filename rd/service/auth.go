package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rd/models"
	"rd/pkg/util"
)

type AuthService struct {
	BaseService
}

func (a *AuthService) Auth (c *gin.Context)  {
	conds := a.getCommonConds()
	conds["username"] = a.getFormString(c, "username")
	conds["password"] = a.getFormString(c, "password")

	user := new(models.Auth).GetOne(conds)
	fmt.Print(user)

	if user == nil {
		a.Failed(c, nil)
		return
	}
	ret := make(map[string]interface{})
	ret["token"], _ = util.GenerateToken(int64(user.ID))
	ret["nickName"] = user.NickName
	a.Success(c, ret)
	return
}

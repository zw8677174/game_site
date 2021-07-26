package api

import (
	"net/http"
	"rd/service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"rd/pkg/app"
	"rd/pkg/e"
	"rd/pkg/util"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func GetAuth(c *gin.Context) {

	valid := validation.Validation{}

	var auth Auth

	if c.BindJSON(&auth) == nil {
	}

	ok, _ := valid.Valid(&auth)

	if !ok {
		app.MarkErrors(valid.Errors)
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := service.AuthService{Username: auth.Username, Password: auth.Password}
	uid := authService.GetUid(c)

	if uid == 0 {
		app.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(uid)
	if err != nil {
		app.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	app.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"token": token,
		"uid": uid,
	})
}

package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"rd/pkg/app"
	"rd/pkg/e"
	"rd/pkg/util"
	"rd/service/auth_service"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func GetAuth(c *gin.Context) {

	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var auth Auth

	if c.BindJSON(&auth) == nil {
	}

	ok, _ := valid.Valid(&auth)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: auth.Username, Password: auth.Password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(auth.Username, auth.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

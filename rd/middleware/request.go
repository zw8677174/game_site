package request

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"rd/pkg/e"
	"rd/pkg/util"
)

type InputMap map[string]interface{}

type PageInfo struct {
	PageNum  int
	PageSize int
}

func GetBaseInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("pageInfo", c.ShouldBindQuery(new(PageInfo)))
		c.Set("body", c.ShouldBindJSON(new(InputMap)))
		c.Set("query", c.ShouldBindQuery(new(InputMap)))
	}
}

// JWT is jwt middleware
func GetParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
			if claims == nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				if claims.Uid > 0 {
					c.Set("uid", claims.Uid)
				}
			}

		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

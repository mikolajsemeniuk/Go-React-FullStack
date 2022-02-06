package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/configuration"
)

func Authorize() gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		cookie, err := context.Request.Cookie("cookie")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
			return
		}

		var token *jwt.Token
		token, err = jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(configuration.Config.GetString("server.secret")), nil
		})

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "not authenticated"})
			return
		}

		context.Set("claims", token.Claims)
	})
}

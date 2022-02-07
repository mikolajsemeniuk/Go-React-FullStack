package middlewares

import (
	"net/http"

	mapset "github.com/deckarep/golang-set"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/configuration"
)

// FIXME:
//	duplicated in `services/account.go`
type Claims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func set(mySlice []string) mapset.Set {
	mySet := mapset.NewSet()
	for _, ele := range mySlice {
		mySet.Add(ele)
	}
	return mySet
}

func Authorize(roles []string) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		cookie, err := context.Request.Cookie("cookie")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
			return
		}

		var token *jwt.Token
		token, err = jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(configuration.Config.GetString("server.secret")), nil
		})

		if err != nil || !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "not authenticated"})
			return
		}

		if len(roles) != 0 {
			s1 := set(token.Claims.(*Claims).Roles)
			s2 := set(roles)
			if !s2.IsSubset(s1) {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"errors":         "not authenticated",
					"current roles":  token.Claims.(*Claims).Roles,
					"required roles": roles,
				})
				return
			}
		}

		context.Set("claims", token.Claims)
	})
}

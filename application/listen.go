package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/configuration"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/controllers"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/inputs"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/middlewares"
)

var (
	router = gin.Default()
)

func Listen() {
	v1 := router.Group(configuration.Config.GetString("server.basepath"))
	{
		auth := v1.Group("auth")
		{
			auth.POST("register", middlewares.Body(inputs.Register{}), controllers.Account.Register)
			auth.POST("login", middlewares.Body(inputs.Login{}), controllers.Account.Login)
			auth.POST("logout", controllers.Account.Logout)
			auth.POST("test", middlewares.Authorize([]string{"admin", "moderator"}), func(c *gin.Context) {
				c.JSON(http.StatusOK, "authenticated")
			})
			auth.POST("test2", middlewares.Authorize([]string{"moderator"}), func(c *gin.Context) {
				c.JSON(http.StatusOK, "authenticated")
			})
			auth.POST("test3", middlewares.Authorize([]string{}), func(c *gin.Context) {
				c.JSON(http.StatusOK, "authenticated")
			})
		}
	}
	router.Run(fmt.Sprintf(":%s", configuration.Config.GetString("server.port")))
}

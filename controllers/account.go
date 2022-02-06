package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/inputs"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/services"
)

var (
	Account IAccount = &account{}
)

type account struct{}

type IAccount interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Logout(*gin.Context)
}

func (*account) Register(context *gin.Context) {
	input := context.MustGet("input").(*inputs.Register)

	token, err := services.Account.Register(input)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.SetCookie("cookie", token, 60*60, "/", "localhost", false, true)
	context.JSON(http.StatusOK, gin.H{
		"message": "user successfully registered in",
	})
}

func (*account) Login(context *gin.Context) {
	input := context.MustGet("input").(*inputs.Login)
	token, err := services.Account.Login(input)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.SetCookie("cookie", token, 60*60, "/", "localhost", false, true)
	context.JSON(http.StatusOK, gin.H{
		"message": "user successfully logged in",
		"token":   token,
	})
}

func (*account) Logout(context *gin.Context) {
	context.SetCookie("cookie", "", -1, "/", "localhost", false, true)

	context.JSON(http.StatusOK, gin.H{
		"message": "user successfully logged out",
	})
}

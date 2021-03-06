package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/configuration"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/data"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/domain"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/inputs"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/repositories"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/sets"
	"golang.org/x/crypto/bcrypt"
)

var (
	Account IAccount = &account{}
)

type account struct{}

type IAccount interface {
	Register(*inputs.Register) (string, error)
	Login(*inputs.Login) (string, error)
}

func (*account) Register(input *inputs.Register) (string, error) {
	var account domain.Account

	// correct
	result := data.Context.Where("username = ?", input.Username).First(&account)
	if result.RowsAffected != 0 {
		return "", errors.New("username already occupied")
	}

	// correct
	result = data.Context.Where("email = ?", input.Email).First(&account)
	if result.RowsAffected != 0 {
		return "", errors.New("email already occupied")
	}

	// correct
	copier.Copy(&account, &input)

	// correct +/-
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		return "", err
	}
	account.Password = password

	// correct
	result = data.Context.Create(&account)
	if result.RowsAffected == 0 {
		return "", errors.New("error has occured")
	}

	// correct +/-
	mapClaims := jwt.MapClaims{}
	mapClaims["Issuer"] = account.Id.String()
	mapClaims["ExpiresAt"] = time.Now().Add(time.Hour).Unix()
	mapClaims["Roles"] = []string{}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	// correct
	token, err := claims.SignedString([]byte(configuration.Config.GetString("server.secret")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (*account) Login(input *inputs.Login) (string, error) {
	// var account domain.Account

	account := *repositories.Account.SingleByEmail(input.Email)

	err := bcrypt.CompareHashAndPassword(account.Password, []byte(input.Password))
	if err != nil {
		return "", errors.New("incorrect password")
	}

	var roles []string
	for _, role := range account.Roles {
		roles = append(roles, role.Role.Name)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, sets.Claims{
		Roles: roles,
		StandardClaims: jwt.StandardClaims{
			Issuer:    account.Id.String(),
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	})

	token, err := claims.SignedString([]byte(configuration.Config.GetString("server.secret")))
	if err != nil {
		return "", err
	}

	return token, nil
}

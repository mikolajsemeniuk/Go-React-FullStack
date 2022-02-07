package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/configuration"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/data"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/domain"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/inputs"
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

	result := data.Context.Where("username = ?", input.Username).First(&account)
	if result.RowsAffected != 0 {
		return "", errors.New("username already occupied")
	}

	result = data.Context.Where("email = ?", input.Email).First(&account)
	if result.RowsAffected != 0 {
		return "", errors.New("email already occupied")
	}

	copier.Copy(&account, &input)

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		return "", err
	}

	account.Password = password
	// account.Roles = []domain.Role{
	// 	{id: 3,}
	// }
	result = data.Context.Create(&account)

	if result.RowsAffected == 0 {
		return "", errors.New("error has occured")
	}

	mapClaims := jwt.MapClaims{}
	mapClaims["Issuer"] = account.Id.String()
	// mapClaims["Issuer"] = strconv.Itoa(account.Id)
	mapClaims["ExpiresAt"] = time.Now().Add(time.Hour).Unix()
	mapClaims["Roles"] = []string{"member"}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	token, err := claims.SignedString([]byte(configuration.Config.GetString("server.secret")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// FIXME:
//	duplicated in `middlewares/authorize.go`
type Claims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func (*account) Login(input *inputs.Login) (string, error) {
	var account domain.Account

	result := data.Context.Where("email = ?", input.Email).Preload("Roles.Role").Take(&account)
	if result.RowsAffected == 0 {
		return "", errors.New("no user with this email address")
	}

	fmt.Println()
	fmt.Println(len(account.Roles))
	fmt.Println()

	err := bcrypt.CompareHashAndPassword(account.Password, []byte(input.Password))
	if err != nil {
		return "", errors.New("incorrect password")
	}

	var roles []string
	for _, role := range account.Roles {
		roles = append(roles, role.Role.Name)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		roles,
		jwt.StandardClaims{
			Issuer: account.Id.String(),
			// Issuer:    strconv.Itoa(account.Id),
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	})

	token, err := claims.SignedString([]byte(configuration.Config.GetString("server.secret")))
	if err != nil {
		return "", err
	}

	return token, nil
}

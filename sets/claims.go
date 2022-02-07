package sets

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

package domain

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

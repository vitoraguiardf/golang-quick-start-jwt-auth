package domain

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId uint   `json:"userId"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

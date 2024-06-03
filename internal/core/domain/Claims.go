package domain

import (
	"github.com/dgrijalva/jwt-go" // upgrade github.com/golang-jwt/jwt/v5
)

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

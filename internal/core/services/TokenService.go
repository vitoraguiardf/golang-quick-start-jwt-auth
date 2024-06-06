package services

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type jwtService struct {
	jwt_secret      string
	jwt_ttl         int
	jwt_refresh_ttl int
}

func NewJWTService() *jwtService {
	return &jwtService{
		jwt_secret:      os.Getenv("JWT_SECRET"),
		jwt_ttl:         60,
		jwt_refresh_ttl: 20160,
	}
}

func (service *jwtService) Login(claims domain.Claims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return accessToken.SignedString([]byte(service.jwt_secret))
}

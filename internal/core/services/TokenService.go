package services

import (
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	jwt_secret      string
	jwt_ttl         int64
	jwt_refresh_ttl int64
}

func NewTokenService() *jwtService {
	instance := &jwtService{
		jwt_secret:      os.Getenv("JWT_SECRET"),
		jwt_ttl:         60,
		jwt_refresh_ttl: 20160,
	}
	if v, e := strconv.Atoi(os.Getenv("JWT_TTL")); e != nil {
		instance.jwt_ttl = int64(v)
	}
	if v, e := strconv.Atoi(os.Getenv("JWT_REFRESH_TTL")); e != nil {
		instance.jwt_refresh_ttl = int64(v)
	}
	return instance
}

func (service *jwtService) NewAccessToken(claims jwt.Claims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return accessToken.SignedString([]byte(service.jwt_secret))
}

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

func (service *jwtService) Create(claims jwt.Claims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return accessToken.SignedString([]byte(service.jwt_secret))
}

func (service *jwtService) Parse(token string) (*jwt.StandardClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(service.jwt_secret), nil
		})
	if err != nil {
		return nil, err
	}
	return parsed.Claims.(*jwt.StandardClaims), nil
}

func (service *jwtService) Validate(token string) bool {
	claims, err := service.Parse(token)
	if err != nil {
		return false
	}
	err = claims.Valid()
	return err == nil
}

func (service *jwtService) Refresh(token string) (string, error) {
	claims, err := service.Parse(token)
	if err != nil {
		return "", err
	}
	err = claims.Valid()
	if err != nil {
		return "", err
	}
	return service.Create(claims)
}

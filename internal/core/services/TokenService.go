package services

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type jwtService struct {
	secret     string
	ttl        int64
	refreshTtl int64
}

func NewTokenService() *jwtService {
	instance := &jwtService{
		secret:     os.Getenv("JWT_SECRET"),
		ttl:        1,
		refreshTtl: 20160,
	}
	if v, e := strconv.Atoi(os.Getenv("JWT_TTL")); e != nil {
		instance.ttl = int64(v)
	}
	if v, e := strconv.Atoi(os.Getenv("JWT_REFRESH_TTL")); e != nil {
		instance.refreshTtl = int64(v)
	}
	return instance
}

func (service *jwtService) Create(claims jwt.Claims) (*domain.Token, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	if token, err := accessToken.SignedString([]byte(service.secret)); err != nil {
		return nil, err
	} else {
		return &domain.Token{
			AccessToken: token,
			TokenType:   "Bearer",
			ExpiresAt:   time.Now().Add(time.Minute * 15).Unix(),
		}, nil
	}
}

func (service *jwtService) Parse(token string) (*jwt.StandardClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(service.secret), nil
		})
	if err != nil {
		return nil, err
	}
	return parsed.Claims.(*jwt.StandardClaims), nil
}

func (service *jwtService) ParseClaims(token string) (*domain.Claims, error) {
	tokenParts := strings.Split(token, "Bearer ")
	parsed, err := jwt.ParseWithClaims(tokenParts[len(tokenParts)-1], &domain.Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(service.secret), nil
		})
	if err != nil {
		return nil, err
	}
	return parsed.Claims.(*domain.Claims), nil
}

func (service *jwtService) Validate(token string) bool {
	claims, err := service.Parse(token)
	if err != nil {
		return false
	}
	err = claims.Valid()
	return err == nil
}

func (service *jwtService) Refresh(token string) (*domain.Token, error) {
	claims, err := service.Parse(token)
	if err != nil {
		return nil, err
	}
	err = claims.Valid()
	if err != nil {
		return nil, err
	}
	return service.Create(*claims)
}

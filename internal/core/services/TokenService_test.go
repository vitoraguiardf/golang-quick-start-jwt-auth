package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
)

var (
	service                                = services.NewTokenService()
	currentAccessToken                     = ""
	currentClaims      *jwt.StandardClaims = nil
)

func TestCreate(t *testing.T) {
	initialClaims := domain.Claims{
		Role: "default", // "default", "user", "operator", "monitor", "supervisor", "admin"
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token, err := service.Create(initialClaims)
	if err != nil {
		t.Fatal(err)
	}
	currentAccessToken = token.AccessToken
}

func TestParse(t *testing.T) {
	claims, err := service.Parse(currentAccessToken)
	if err != nil {
		t.Fatal(err)
	}
	currentClaims = claims
}

func TestValidate(t *testing.T) {
	if !service.Validate(currentAccessToken) {
		t.Fatal(errors.New("Invalid Token"))
	}
}

func TestRefresh(t *testing.T) {
	token, err := service.Refresh(currentAccessToken)
	if err != nil {
		t.Fatal(err)
	}
	currentAccessToken = token.AccessToken
}

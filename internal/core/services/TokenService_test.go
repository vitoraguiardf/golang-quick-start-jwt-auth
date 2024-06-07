package services_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
)

var service = services.NewTokenService()

func TestNewAccessToken(t *testing.T) {
	claims := domain.Claims{
		Roles: []string{"default", "user", "operator", "monitor", "supervisor", "admin"},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	_, err := service.NewAccessToken(claims)
	if err != nil {
		claimErr := fmt.Errorf("Fail for claims %v", claims)
		t.Fatal(errors.Join(claimErr, err))
	}
}

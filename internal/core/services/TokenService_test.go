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

func TestCreateToken(t *testing.T) {
	roleList := []string{"default", "user", "operator", "monitor", "supervisor", "admin"}
	for _, role := range roleList {
		createTokenRoleTest(t, role)
	}
}

func createTokenRoleTest(t *testing.T, role string) error {
	fmt.Printf("Running tests with %v Role\n", role)
	claims := domain.Claims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	fmt.Printf("\tClaims %v\n", claims)
	_, err := service.NewAccessToken(claims)
	if err != nil {
		claimErr := fmt.Errorf("Fail for claims %v", claims)
		t.Fatal(errors.Join(claimErr, err))
	}
	return nil
}

package services_test

import (
	"testing"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
)

func TestLogin(t *testing.T) {
	service := services.NewJWTService()
	token, err := service.Login(domain.Claims{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

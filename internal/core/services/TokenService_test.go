package services_test

import (
	"testing"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
)

var service = services.NewJWTService()

func TestLogin(t *testing.T) {
	token, err := service.Login(domain.Claims{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestMe(t *testing.T) {
	token, err := service.Me(domain.Claims{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestRefresh(t *testing.T) {
	token, err := service.Refresh(domain.Claims{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestLogout(t *testing.T) {
	token, err := service.Logout(domain.Claims{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

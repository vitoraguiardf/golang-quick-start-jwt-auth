package services_test

import (
	"testing"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
)

var authService = services.NewAuthService()

func TestAuthServiceLogin(t *testing.T) {
	_, err := authService.Login()
	if err != nil {
		t.Error(err)
	}
}

func TestAuthServiceMe(t *testing.T) {
	_, err := authService.Me()
	if err != nil {
		t.Error(err)
	}
}

func TestAuthServiceRefresh(t *testing.T) {
	_, err := authService.Refresh()
	if err != nil {
		t.Error(err)
	}
}

func TestAuthServiceLogout(t *testing.T) {
	_, err := authService.Logout()
	if err != nil {
		t.Error(err)
	}
}

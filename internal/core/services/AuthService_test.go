package services_test

import (
	"testing"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/services"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/repositories/authrepository"
)

var authRepository = authrepository.NewSqliteRepository()
var authService = services.NewAuthService(authRepository)
var authorization domain.Token
var credentials = domain.Credentials{Email: "vitoraguiar.df@gmail.com", Password: "password@123"}

func TestAuthServiceLogin(t *testing.T) {
	tk, err := authService.Login(credentials)
	if err != nil {
		t.Error(err)
	}
	authorization = *tk
}

func TestAuthServiceMe(t *testing.T) {
	_, err := authService.Me(authorization.AccessToken)
	if err != nil {
		t.Error(err)
	}
}

func TestAuthServiceRefresh(t *testing.T) {
	_, err := authService.Refresh(authorization.AccessToken)
	if err != nil {
		t.Error(err)
	}
}

func TestAuthServiceLogout(t *testing.T) {
	_, err := authService.Logout(authorization.AccessToken)
	if err != nil {
		t.Error(err)
	}
}

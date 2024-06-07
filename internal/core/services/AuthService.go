package services

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type authService struct {
	repository ports.AuthRepository
}

func NewAuthService() *authService {
	return &authService{}
}

func (service *authService) Login() (string, error) {
	return "Successful login!", nil
}

func (service *authService) Me() (domain.User, error) {
	return domain.User{}, nil
}

func (service *authService) Refresh() (string, error) {
	return "Successful refresh!", nil
}

func (service *authService) Logout() (string, error) {
	return "Successful logout!", nil
}

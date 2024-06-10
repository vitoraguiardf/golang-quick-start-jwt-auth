package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type authService struct {
	repository ports.AuthRepository
	jwtService jwtService
}

func NewAuthService(authRepository ports.AuthRepository) *authService {
	return &authService{
		repository: authRepository,
		jwtService: *NewTokenService(),
	}
}

func (service *authService) Login(username string, password string) (string, error) {
	// TODO Login - credentials validation
	if username != "vitoraguiardf" {
		return "", errors.New("user not found")
	}
	if password != "123" {
		return "", errors.New("incorrect password")
	}
	claims := domain.Claims{
		Roles: []string{"default", "user", "operator", "monitor", "supervisor", "admin"},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	//
	if token, err := service.jwtService.Create(claims); err != nil {
		return "", err
	} else {
		return token, nil
	}
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

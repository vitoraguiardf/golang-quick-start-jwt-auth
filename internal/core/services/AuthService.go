package services

import (
	"errors"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/utils"
)

type authService struct {
	repository   ports.AuthRepository
	tokenService jwtService
}

func NewAuthService(authRepository ports.AuthRepository) *authService {
	return &authService{
		repository:   authRepository,
		tokenService: *NewTokenService(),
	}
}

func (s *authService) Login(credentials domain.Credentials) (string, error) {
	var user *domain.User = nil
	if !s.repository.ExistsByEmail(credentials.Email) {
		return "", errors.New("user not found")
	}
	if r, err := s.repository.FindUserByEmail(credentials.Email); err != nil {
		return "", err
	} else {
		user = r
	}
	if user == nil {
		return "", errors.New("auth service has failed")
	}
	if err := checkCredentials(*user, credentials); err != nil {
		return "", err
	}
	if token, err := s.tokenService.Create(user.Claims()); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func (s *authService) Me() (domain.User, error) {
	return domain.User{}, nil
}

func (s *authService) Refresh() (string, error) {
	return "Successful refresh!", nil
}

func (s *authService) Logout() (string, error) {
	return "Successful logout!", nil
}

func checkCredentials(u domain.User, c domain.Credentials) error {
	if u.Email != c.Email {
		return errors.New("invalid email adress")
	}
	if err := utils.HashCompare(u.Password, c.Password); err != nil {
		return errors.New("invalid password")
	}
	return nil
}

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

func (s *authService) Login(credentials domain.Credentials) (*domain.Token, error) {
	var user *domain.User = nil
	if !s.repository.ExistsByEmail(credentials.Email) {
		return nil, errors.New("user not found")
	}
	if r, err := s.repository.FindUserByEmail(credentials.Email); err != nil {
		return nil, err
	} else {
		user = r
	}
	if user == nil {
		return nil, errors.New("auth service has failed")
	}
	if err := checkCredentials(*user, credentials); err != nil {
		return nil, err
	}
	if authorization, err := s.tokenService.Create(user.Claims()); err != nil {
		return nil, err
	} else {
		return authorization, nil
	}
}

func (s *authService) Me(authorization string) (*domain.User, error) {
	claims, err := s.tokenService.ParseClaims(authorization)
	if err != nil {
		return nil, err
	}
	if err := claims.Valid(); err != nil {
		return nil, err
	}
	user, err := s.repository.FindUser(claims.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) Refresh(authorization string) (*domain.Token, error) {
	return s.tokenService.Refresh(authorization)
}

func (s *authService) Logout(authorization string) (string, error) {
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

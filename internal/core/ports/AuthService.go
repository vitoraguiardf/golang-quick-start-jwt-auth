package ports

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type AuthService interface {
	Login(domain.Credentials) (token *domain.Token, err error)
	Me(authorization string) (*domain.User, error)
	Refresh(authorization string) (token *domain.Token, err error)
	Logout(authorization string) (message string, err error)
}

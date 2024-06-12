package ports

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type AuthService interface {
	Login(domain.Credentials) (string, error)
	Me() (domain.User, error)
	Refresh() (string, error)
	Logout() (string, error)
}

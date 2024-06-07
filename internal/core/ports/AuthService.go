package ports

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type AuthService interface {
	Login(username string, password string) (string, error)
	Me() (domain.User, error)
	Refresh() (string, error)
	Logout() (string, error)
}

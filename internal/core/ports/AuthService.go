package ports

import (
	"github.com/golang-jwt/jwt"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type AuthService interface {
	Login() (jwt.Token, error)
	Me() (domain.User, error)
	Refresh() (jwt.Token, error)
	Logout() (string, error)
}

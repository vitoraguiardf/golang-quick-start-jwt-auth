package ports

import "github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"

type UserService interface {
	Get(id uint) (domain.User, error)
}

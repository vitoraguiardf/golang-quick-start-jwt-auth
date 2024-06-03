package ports

import "github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"

type UserRepository interface {
	Get(id uint) (domain.User, error)
	Save(domain.User) error
}

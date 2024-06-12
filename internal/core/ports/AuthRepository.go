package ports

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type AuthRepository interface {
	ExistsByEmail(email string) (exists bool)
	FindUserByEmail(email string) (user *domain.User, err error)
}

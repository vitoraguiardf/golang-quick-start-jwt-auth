package ports

import "github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"

type UserService interface {
	FindAll() (items map[string]domain.User, err error)
	FindById(id string) (item *domain.User, err error)
	Update(id string, model *domain.User) (err error)
	Replace(id string, model *domain.User) (err error)
	Create(model *domain.User) (err error)
	Delete(id string) (err error)
}

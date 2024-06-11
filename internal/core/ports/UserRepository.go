package ports

import "github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"

type UserRepository interface {
	FindAll() (items *[]domain.User, err error)
	FindById(id uint) (item *domain.User, err error)
	Update(id uint, model *domain.User) (err error)
	Replace(id uint, model *domain.User) (err error)
	Create(model *domain.User) (err error)
	Delete(id uint) (err error)
}

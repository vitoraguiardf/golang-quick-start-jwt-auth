package services

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type userService struct {
	repository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *userService {
	return &userService{
		repository: userRepository,
	}
}

func (service *userService) Get(id uint) (domain.User, error) {
	user, err := service.repository.Get(id)
	if err != nil {
		return domain.User{}, err // get model from repository has failed
	}
	return user, nil
}

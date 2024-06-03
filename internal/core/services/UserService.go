package services

import (
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/ports"
)

type service struct {
	repository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *service {
	return &service{
		repository: userRepository,
	}
}

func (svc *service) Get(id uint) (domain.User, error) {
	user, err := svc.repository.Get(id)
	if err != nil {
		return domain.User{}, err // get model from repository has failed
	}
	return user, nil
}

package services

import (
	"errors"
	"regexp"
	"strconv"

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

func (service *userService) Create(user *domain.User) error {
	if err := service.repository.Create(user); err != nil {
		return err
	}
	return nil
}

func (service *userService) FindAll() (map[string]domain.User, error) {
	if items, err := service.repository.FindAll(); err != nil {
		return nil, err
	} else {
		return items, nil
	}
}

func (service *userService) FindById(idParam string) (*domain.User, error) {
	if id, err := parseIdParam(idParam); err != nil {
		return nil, err
	} else {
		if item, err := service.repository.FindById(id); err != nil {
			return nil, err
		} else {
			return item, nil
		}
	}
}

func (service *userService) Update(idParam string, model *domain.User) error {
	if id, err := parseIdParam(idParam); err != nil {
		return err
	} else {
		if err := service.repository.Update(id, model); err != nil {
			return err
		}
	}
	return nil
}

func (service *userService) Replace(idParam string, model *domain.User) error {
	if id, err := parseIdParam(idParam); err != nil {
		return err
	} else {
		if err := service.repository.Replace(id, model); err != nil {
			return err
		}
	}
	return nil
}

func (service *userService) Delete(idParam string) error {
	if id, err := parseIdParam(idParam); err != nil {
		return err
	} else {
		if err := service.repository.Delete(id); err != nil {
			return err
		}
	}
	return nil
}

func parseIdParam(idParam string) (id uint, err error) {
	if match, err := regexp.MatchString("^[0-9]+$", idParam); err != nil {
		return 0, err
	} else if !match {
		return 0, errors.New("invalid value for id param")
	}
	if val, err := strconv.Atoi(idParam); err != nil {
		return 0, err
	} else {
		return uint(val), nil
	}
}

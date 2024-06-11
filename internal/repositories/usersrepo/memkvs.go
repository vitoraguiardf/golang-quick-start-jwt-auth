package usersrepo

import (
	"encoding/json"
	"errors"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type userRepository struct {
	kvs map[uint][]byte
}

func NewKvsRepository() *userRepository {
	return &userRepository{kvs: map[uint][]byte{}}
}

func (repository *userRepository) Get(id uint) (domain.User, error) {
	if value, ok := repository.kvs[id]; ok {
		user := domain.User{}
		err := json.Unmarshal(value, &user)
		if err != nil {
			return domain.User{}, errors.New("fail to get value from kvs")
		}
		return user, nil
	}
	return domain.User{}, errors.New("user not found")
}

func (repostory *userRepository) Save(user domain.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return errors.New("users fails at marshal into json string")
	}
	repostory.kvs[user.ID] = bytes
	return nil
}

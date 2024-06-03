package usersrepository

import (
	"encoding/json"
	"errors"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
)

type userRepository struct {
	kvs map[uint][]byte
}

func NewRepository() *userRepository {
	return &userRepository{kvs: map[uint][]byte{}}
}

func (repository *userRepository) Get(id uint) (domain.User, error) {
	return domain.User{}, nil
}

func (repostory *userRepository) Save(user domain.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return errors.New("users fails at marshal into json string")
	}
	repostory.kvs[user.ID] = bytes
	return nil
}

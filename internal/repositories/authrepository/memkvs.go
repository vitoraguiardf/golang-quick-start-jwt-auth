package authrepository

import (
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt"
)

type kvsRepository struct {
	kvs map[uint][]byte
}

func NewKvsRepository() *kvsRepository {
	return &kvsRepository{
		kvs: map[uint][]byte{},
	}
}

func (repository *kvsRepository) Get(id uint) (jwt.Token, error) {
	if value, ok := repository.kvs[id]; ok {
		item := jwt.Token{}
		err := json.Unmarshal(value, &item)
		if err != nil {
			return jwt.Token{}, errors.New("fail to get value from kvs")
		}
		return item, nil
	}
	return jwt.Token{}, nil
}

func (repostory *kvsRepository) Save(item jwt.Token) error {
	bytes, err := json.Marshal(item)
	if err != nil {
		return errors.New("users fails at marshal into json string")
	}
	repostory.kvs[1] = bytes
	return nil
}

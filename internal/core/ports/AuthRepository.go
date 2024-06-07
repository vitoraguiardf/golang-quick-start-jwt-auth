package ports

import "github.com/golang-jwt/jwt"

type AuthRepository interface {
	Get(id uint) (jwt.Token, error)
	Save(jwt.Token) error
}

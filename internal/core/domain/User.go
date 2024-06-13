package domain

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
	Name      string       `json:"name" binding:"required"`
	Email     string       `json:"email" gorm:"unique" binding:"required"`
	Password  string       `json:"-" binding:"required"`
	Role      string       `json:"role"`
}

func (u *User) Claims() *Claims {
	return &Claims{
		UserId: u.ID,
		Role:   u.Role, // "default", "user", "operator", "monitor", "supervisor", "admin"
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
}

package authrepository

import (
	"errors"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/database/sqlite"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	DB *gorm.DB
}

func NewSqliteRepository() *sqliteRepository {
	return &sqliteRepository{
		DB: sqlite.SqliteDB,
	}
}

func (repo *sqliteRepository) ExistsByEmail(email string) bool {
	var user *domain.User
	res := repo.DB.Where("Email = ?", email).First(&user)
	return res.RowsAffected == 1 && user != nil
}

func (repo *sqliteRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user *domain.User
	if res := repo.DB.Where("Email = ?", email).First(&user); res.RowsAffected != 1 || user == nil {
		return nil, errors.New("auth service has failed")
	}
	return user, nil
}

func (repo *sqliteRepository) FindUser(id uint) (*domain.User, error) {
	var user domain.User
	repo.DB.Where("ID = ?", id).First(&user)
	return &user, nil
}

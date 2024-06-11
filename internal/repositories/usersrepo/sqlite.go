package usersrepo

import (
	"fmt"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/repositories/persistence"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	DB *gorm.DB
}

func NewSqliteRepository() *sqliteRepository {
	return &sqliteRepository{
		DB: persistence.SqliteDB,
	}
}

func (repo *sqliteRepository) Get(id uint) (domain.User, error) {
	var user domain.User
	repo.DB.Where("ID = ?", id).First(&user)
	return user, nil
}

func (repo *sqliteRepository) Save(user domain.User) error {
	res := repo.DB.Save(user)
	affected := res.RowsAffected
	if affected == 1 {
		return nil
	}
	return fmt.Errorf("%v rows affected when expexted only one", affected)
}

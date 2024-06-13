package usersrepo

import (
	"fmt"

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

func (repo *sqliteRepository) Create(user *domain.User) (err error) {
	res := repo.DB.Create(user)
	affected := res.RowsAffected
	if affected == 1 {
		return nil
	}
	return fmt.Errorf("%v rows affected when expected only one", affected)
}

func (repo *sqliteRepository) FindAll() (map[string]domain.User, error) {
	items := map[string]domain.User{}
	repo.DB.Find(&items)
	return items, nil
}

func (repo *sqliteRepository) FindById(id uint) (*domain.User, error) {
	var user domain.User
	repo.DB.Where("ID = ?", id).First(&user)
	return &user, nil
}

func (repo *sqliteRepository) Update(id uint, model *domain.User) (err error) {
	return nil
}

func (repo *sqliteRepository) Replace(id uint, model *domain.User) (err error) {
	return nil
}

func (repo *sqliteRepository) Delete(id uint) (err error) {
	return nil
}

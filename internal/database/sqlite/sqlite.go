package sqlite

import (
	"os"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/core/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDB struct {
	name string
}

func NewSqliteDB() *sqliteDB {
	var _name string
	_name = os.Getenv("DB_DATABASE")
	if len(_name) >= 0 {
		_name = os.Getenv("APP_NAME")
		if len(_name) >= 0 {
			_name = "golang"
		}
	}
	return &sqliteDB{
		name: _name,
	}
}

var SqliteDB *gorm.DB

func (c *sqliteDB) InitDB() error {
	db, err := gorm.Open(sqlite.Open(c.name+".db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&domain.User{},
		&domain.Claims{},
	); err != nil {
		return err
	}

	SqliteDB = db
	return nil
}

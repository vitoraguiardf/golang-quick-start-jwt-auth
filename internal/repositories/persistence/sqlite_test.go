package persistence_test

import (
	"testing"

	"github.com/vitoraguiardf/golang-quick-start-jwt-auth/internal/repositories/persistence"
)

func TestConnection(t *testing.T) {
	t.Setenv("DB_DATABASE", "test.db")
	persistence.NewSqliteDB().InitDB()
}

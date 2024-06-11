package persistence

import (
	"testing"
)

func TestConnection(t *testing.T) {
	t.Setenv("DB_DATABASE", "test.db")
	NewSqliteDB().InitDB()
}

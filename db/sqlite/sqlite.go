package sqlite

import (
	"fmt"

	"github.com/CloudSilk/pkg/db"
	gormsqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewSqlite(connStr string, debug bool) db.DBClientInterface {
	dbClient, err := gorm.Open(gormsqlite.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db.NewDBClient(dbClient, debug)
}

// NewSqlite2 creates a SQLite client using glebarez/go-sqlite (pure Go, no CGO).
// Note: _loc, _auth parameters are mattn/go-sqlite3 (CGO) specific and NOT supported here.
// Use _pragma for PRAGMA settings instead.
func NewSqlite2(user, password, path, dbName string, debug bool) db.DBClientInterface {
	return NewSqlite(fmt.Sprintf("%s?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)", path), debug)
}

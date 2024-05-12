package sqlite

import (
	"fmt"

	"github.com/CloudSilk/pkg/db"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// NewSqlite New Sqlite
func NewSqlite(connStr string, debug bool) db.DBClientInterface {
	dbClient, err := gorm.Open(sqlite.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = dbClient.Exec("PRAGMA encoding='UTF-8'").Error
	if err != nil {
		panic(err)
	}
	return db.NewDBClient(dbClient, debug)
}

// NewSqlite2 New Sqlite
// sqlite权限访问就是个鸡肋，没啥用，只能防君子不能防小人
func NewSqlite2(user, password, path, dbName string, debug bool) db.DBClientInterface {
	return NewSqlite(fmt.Sprintf("%s?cache=private&cache_size=1024&mode=rwc&_loc=Asia%%2FShanghai&_auth&_auth_user=%s&_auth_pass=%s", path, user, password), debug)
}

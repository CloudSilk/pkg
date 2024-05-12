package sqlserver

import (
	"fmt"

	"github.com/CloudSilk/pkg/db"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewSqlServer New SqlServer
func NewSqlServer(connStr string, debug bool) db.DBClientInterface {
	dbClient, err := gorm.Open(sqlserver.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db.NewDBClient(dbClient, debug)
}

// NewSqlServer2 New SqlServer
func NewSqlServer2(userName, password, host, dbName string, port int, debug bool) db.DBClientInterface {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", userName, password, host, port, dbName)
	return NewSqlServer(dsn, debug)
}

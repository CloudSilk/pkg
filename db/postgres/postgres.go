package postgres

import (
	"fmt"

	"github.com/CloudSilk/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgres New Postgres
func NewPostgres(connStr string, debug bool) db.DBClientInterface {
	dbClient, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	return db.NewDBClient(dbClient, debug)
}

// NewPostgres2 New Postgres
func NewPostgres2(userName, password, host, dbName string, port int, debug bool) db.DBClientInterface {
	return NewPostgres(fmt.Sprintf("postgresql://%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, host, port, dbName), debug)
}

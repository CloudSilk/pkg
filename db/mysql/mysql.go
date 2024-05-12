package mysql

import (
	"fmt"

	"github.com/CloudSilk/pkg/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMysql New Mysql
func NewMysql(connStr string, debug bool) db.DBClientInterface {
	dbClient, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db.NewDBClient(dbClient, debug)
}

// NewMysql2 New Mysql
func NewMysql2(userName, password, host, dbName string, port int, debug bool) db.DBClientInterface {
	return NewMysql(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, port, dbName), debug)
}

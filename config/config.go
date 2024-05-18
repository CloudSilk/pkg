package config

import (
	"os"

	"github.com/CloudSilk/pkg/db"
	"github.com/CloudSilk/pkg/db/mysql"
	"github.com/CloudSilk/pkg/db/postgres"
	"github.com/CloudSilk/pkg/db/sqlite"
	"gopkg.in/yaml.v3"
)

var DefaultConfig = &Config{}

func InitFromFile(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, DefaultConfig)
}

func InitFromString(str string) error {
	return yaml.Unmarshal([]byte(str), DefaultConfig)
}

func GetDBConfig(name string) (*DBConfig, bool) {
	dbConfig, ok := DefaultConfig.DBConfigs[name]
	return dbConfig, ok
}

func NewDB(name string) (bool, db.DBClientInterface) {
	dbConfig, ok := GetDBConfig(name)
	if !ok {
		return false, nil
	}
	switch dbConfig.DBType {
	case "mysql":
		if dbConfig.ConnectionStr != "" {
			return true, mysql.NewMysql(dbConfig.ConnectionStr, DefaultConfig.Debug)
		} else {
			return true, mysql.NewMysql2(dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.DBName, dbConfig.Port, DefaultConfig.Debug)
		}
	case "sqlite":
		if dbConfig.ConnectionStr != "" {
			sqlite.NewSqlite(dbConfig.ConnectionStr, DefaultConfig.Debug)
		} else {
			return true, sqlite.NewSqlite2(dbConfig.UserName, dbConfig.Password, dbConfig.FileName, dbConfig.DBName, DefaultConfig.Debug)
		}
	case "postgres":
		if dbConfig.ConnectionStr != "" {
			return true, postgres.NewPostgres(dbConfig.ConnectionStr, DefaultConfig.Debug)
		} else {
			return true, postgres.NewPostgres2(dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.DBName, dbConfig.Port, DefaultConfig.Debug)
		}
	}
	return false, nil
}

type Config struct {
	DBConfigs        map[string]*DBConfig `yaml:"dbConfigs"`
	Debug            bool                 `yaml:"debug"`
	Token            TokenConfig          `yaml:"token"`
	SuperAdminRoleID string               `yaml:"superAdminRoleID"`
	PlatformTenantID string               `yaml:"platformTenantID"`
	DefaultRoleID    string               `yaml:"defaultRoleID"`
	DefaultPwd       string               `yaml:"defaultPwd"`
	MiniApp          []MiniAppConfig      `yaml:"miniApp"`
	EnableTenant     bool                 `yaml:"enableTenant"`
}

type DBConfig struct {
	DBType        string `yaml:"dbType"`
	ConnectionStr string `yaml:"connectionStr"`
	FileName      string `yaml:"fileName"`
	UserName      string `yaml:"userName"`
	Password      string `yaml:"password"`
	Host          string `yaml:"Host"`
	DBName        string `yaml:"dbName"`
	Port          int    `yaml:"port"`
}

type MiniAppConfig struct {
	ID       string `yaml:"id"`
	Name     string `yaml:"name"`
	Secret   string `yaml:"secret"`
	TenantID string `yaml:"tenantID"`
}

type TokenConfig struct {
	Key       string `yaml:"key"`
	RedisAddr string `yaml:"redisAddr"`
	RedisName string `yaml:"redisName"`
	RedisPwd  string `yaml:"redisPwd"`
	Expired   int    `yaml:"expired"`
}

package repo

import (
	"os"

	"github.com/joho/godotenv" // импортируем пакет
)

const (
	dbUri      = "DB_URI"
	dbaddr     = "DB_ADDR"
	dbUsername = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"
)

var vars Config

func init() {
	_ = godotenv.Load()

	vars = newConfig()
}

func GetVars() Config {
	return vars
}

type Config interface {
	GetDbUri() string
	GetDBaddr() string
	GetDbUsername() string
	GetDbPasswd() string
}

type config struct {
	dbUri      string
	dbaddr     string
	dbUsername string
	dbPassword string
}

func newConfig() Config {
	conf := new(config)
	conf.SetDbUri()
	conf.SetDBAddr()
	conf.SetDbName()
	conf.SetDbPasswd()

	return conf
}

func (conf *config) GetDbUri() string { return conf.dbUri }
func (conf *config) SetDbUri()        { conf.dbUri = mustGetenv(dbUri) }

func (conf *config) GetDBaddr() string { return conf.dbaddr }
func (conf *config) SetDBAddr()        { conf.dbaddr = mustGetenv(dbaddr) }

func (conf *config) GetDbUsername() string { return conf.dbUsername }
func (conf *config) SetDbName()            { conf.dbUsername = mustGetenv(dbUsername) }

func (conf *config) GetDbPasswd() string { return conf.dbPassword }
func (conf *config) SetDbPasswd()        { conf.dbPassword = mustGetenv(dbPassword) }

func mustGetenv(key string) string {
	if v := os.Getenv(key); len(v) >= 0 {
		return v
	}

	panic(key)
}

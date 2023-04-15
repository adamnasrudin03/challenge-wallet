package configs

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name      string
	Env       string
	Port      string
	SecretKey string
}

type DbConfig struct {
	Host        string
	Port        string
	Dbname      string
	Username    string
	Password    string
	DbIsMigrate bool
	DebugMode   bool
}

type Configs struct {
	Appconfig AppConfig
	Dbconfig  DbConfig
}

var lock = &sync.Mutex{}
var configs *Configs

func GetInstance() *Configs {
	if configs == nil {
		lock.Lock()

		if err := godotenv.Load(); err != nil {
			log.Println("Failed to load env file")
		}

		configs = &Configs{
			Appconfig: AppConfig{
				Name:      getEnv("APP_NAME", "to-do-list"),
				Env:       getEnv("APP_ENV", "dev"),
				Port:      getEnv("APP_PORT", "3030"),
				SecretKey: getEnv("SECRET_KEY", "SECRET-KEY"),
			},

			Dbconfig: DbConfig{
				Host:        getEnv("MYSQL_HOST", "localhost"),
				Port:        getEnv("MYSQL_PORT", "3306"),
				Username:    getEnv("MYSQL_USER", "root"),
				Password:    getEnv("MYSQL_PASSWORD", ""),
				Dbname:      getEnv("MYSQL_DBNAME", "test_db"),
				DbIsMigrate: getEnv("DB_ISMIGRATE", "true") == "true",
				DebugMode:   getEnv("DEBUG_MODE", "true") == "true",
			},
		}
		lock.Unlock()
	}

	return configs
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

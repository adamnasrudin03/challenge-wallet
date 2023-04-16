package database

import (
	"fmt"
	"log"

	"adamnasrudin03/challenge-wallet/app/configs"
	"adamnasrudin03/challenge-wallet/app/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Setup Db Connection is creating a new connection to our database
func SetupMySQLConnection() *gorm.DB {
	configs := configs.GetInstance()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configs.Dbconfig.Username,
		configs.Dbconfig.Password,
		configs.Dbconfig.Host,
		configs.Dbconfig.Port,
		configs.Dbconfig.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	if configs.Dbconfig.DebugMode {
		db = db.Debug()
	}

	if configs.Dbconfig.DbIsMigrate {
		//auto migration entity db
		db.AutoMigrate(
			entity.User{},
			entity.Wallet{},
			entity.Transaction{},
		)
	}

	log.Println("Connection Database Success!")
	return db
}

// Close Db Connection method is closing a connection between your app and your db
func CloseMySQLConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}

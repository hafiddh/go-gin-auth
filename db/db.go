package db

import (
	"log"
	"os"

	"github.com/hafiddh/go-gin-auth/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBcon *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("dsn")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("DB Conn Err")
	}
	log.Println("DB Conn Successful")

	db.AutoMigrate(new(model.User))

	DBcon = db
}

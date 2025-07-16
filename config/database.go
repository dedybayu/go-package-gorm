package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connectionString := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local",
		ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connection established successfully")
}


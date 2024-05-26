package configs

import (
	"fmt"
	"go-restapi-fiber-lem/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectDB(config Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helpers.ErrorPanic(err)

	log.Println("Connected database success...")
	return db
}

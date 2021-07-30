package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GetConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}

package datastore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDB() *gorm.DB {
	dsn := "root:****@tcp(localhost:3306)/****?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

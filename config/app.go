package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open("mysql", "karma:tennyson@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("failed to connect to database!")
	}
}

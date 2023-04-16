package models

import (
	"github.com/ankan792/bookmanagementsystem-GO/config"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

// initialise database
func init() {
	config.ConnectToDB()
	config.DB.AutoMigrate(&Book{})
	db = config.DB
}

func Create(b *Book) *Book {
	if db.NewRecord(b) {
		db.Create(&b)
	}
	return b
}

func GetAll() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetByID(id int) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return &book, result.Error
	}
	return &book, nil

}

func Update(b *Book) {
	db.Save(&b)
}

func DeleteByID(id int) {
	db.Delete(&Book{}, id)
}

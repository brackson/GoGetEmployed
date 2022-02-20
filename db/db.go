package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	// "GoGetEmployed/config"
)

var db *gorm.DB

func Init() {
	// c := config.GetConfig()

	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// // Creating the table
	// if !db.HasTable(&Users{}) {
	// 	db.CreateTable(&Users{})
	// 	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	// }
}

func GetDB() *gorm.DB {
	return db
}

package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func CreateConnection() error {
	connectionString := "root:j1751502021@tcp(database:3306)/notesAPI?parseTime=true"
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("connection made successfully ")
	db = d
	return nil
}
func GetConnection() *gorm.DB {
	return db
}

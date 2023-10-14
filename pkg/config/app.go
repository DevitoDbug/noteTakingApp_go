package config

import "github.com/jinzhu/gorm"

var db *gorm.DB

func CreateConnection() error {
	connectionString := ""
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	db = d
	return nil
}
func GetConnection() *gorm.DB {
	return db
}

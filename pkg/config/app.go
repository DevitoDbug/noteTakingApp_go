package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func CreateConnection() error {
	connectionString := "root:j1751502021@tcp(database:3306)/noteAPIDatabase?parseTime=true"

	// Retry the connection until the database is available
	for {
		d, err := gorm.Open("mysql", connectionString)
		if err != nil {
			fmt.Printf("Error connecting to the database: %v. Retrying...\n", err)
			time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
			continue
		}

		fmt.Println("Connection made successfully")
		db = d
		return nil
	}
}

func GetConnection() *gorm.DB {
	return db
}

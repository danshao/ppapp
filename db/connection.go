package db

import (
	"log"
	"sync"
	_ "sync"

	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DBConnection *gorm.DB
var once sync.Once

// SharedConnection acts as a DB connection singleton
func SharedConnection() *gorm.DB {
	once.Do(func() {
		var err error
		DBConnection, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/pp")

		if err != nil {
			log.Println("Error connecting to MySQL DB: %v", err)
		}

		DBConnection.DB().SetMaxIdleConns(5)
		DBConnection.DB().SetMaxOpenConns(50)
		DBConnection.DB().SetConnMaxLifetime(30 * time.Second)

		DBConnection.LogMode(false)

		return
	})

	return DBConnection
}

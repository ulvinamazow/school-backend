package config

import (
	"fmt"
	"time"

	"github.com/ulvinamazow/studentAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnnectDatabase() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:school_db@localhost/ulvinamazow?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for sqlDB.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect database")
			break
		}
	}
	fmt.Println("Connected to database: ", sqlDB.Ping() == nil)
	db.AutoMigrate(&models.User{})
}

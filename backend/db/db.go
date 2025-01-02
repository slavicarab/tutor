package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"tutor/backend/models"
)

func InitDB() *gorm.DB {
	// Replace with your database credentials
	dsn := "tutor:tutor@tcp(127.0.0.1:3306)/tutor?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate models
	db.AutoMigrate(&models.Users{}, &models.Appointment{})
	log.Println("Database connection established and models migrated")
	return db
}

package db

import (
	"log"
	"os"
	"romantic-situation/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// マイグレーション
	err = DB.AutoMigrate(
		&models.User{},
		&models.Situation{},
		&models.Favorite{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
} 
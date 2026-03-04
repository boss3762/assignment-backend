package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
	"os"
	"github.com/joho/godotenv"
	"log/slog"
	"fmt"
)

var DB *gorm.DB

func ConnectDB() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	err := godotenv.Load()
	if err != nil {
		logger.Error("Failed to load environment variables", "error", err)
	}
	// db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))
	// if err != nil {
	// 	logger.Error("Failed to connect to database", "error", err)
	// 	os.Exit(1)
	// }

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection successfully opened")
	DB = db

	
}

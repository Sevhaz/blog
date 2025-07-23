package database

import (
	"blog/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDb() *gorm.DB {
	err := godotenv.Load()
	if err !=nil{
		log.Fatal("Error Loading Env", err)
	}
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	Db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = Db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		log.Fatal("Failed to enable uuid-ossp extension:", err)
	}

	err = Db.AutoMigrate(&models.User{}, &models.Blog{})

	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

	fmt.Println("Connected to database successfully")
	return Db
}

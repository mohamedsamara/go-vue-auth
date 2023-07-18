package db

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/mohamedsamara/golang-vue/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbURL := os.Getenv("POSTGRES_DATABASE_URL")
	if dbURL == "" {
		host := os.Getenv("POSTGRES_HOST")
		dbPort := os.Getenv("POSTGRES_PORT")
		dbName := os.Getenv("POSTGRES_DB")
		dbUsername := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, dbUsername, dbPassword, dbName, dbPort)
	}

	var err error

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURL,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("🚀 db connected")

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.AutoMigrate(&models.User{})

	fmt.Println("👍 Migration complete.")

	return db
}

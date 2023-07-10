package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type database struct {
	db *gorm.DB
}

var DB database

func InitDB() {
	dbURL := os.Getenv("POSTGRES_DATABASE_URL")
	if dbURL == "" {
		host := os.Getenv("POSTGRES_HOST")
		dbPort := os.Getenv("POSTGRES_PORT")
		dbName := os.Getenv("POSTGRES_DB")
		dbUsername := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, dbUsername, dbPassword, dbName, dbPort)
	}

	if dbURL == "" {
		panic("DB env vars were not found")
	}
	var err error
	db, err := gorm.Open("postgres", dbURL)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	DB.db = db

	fmt.Println("db connected", dbURL)
}

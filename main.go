package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/mohamedsamara/golang-vue/db"
	"github.com/mohamedsamara/golang-vue/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("main: no .env file")
	}

	db.InitDB()

	router := routes.InitRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	log.Println("Listening on PORT :", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

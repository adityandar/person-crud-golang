package main

import (
	"database/sql"
	"fmt"
	"mini_project_restapi/database"
	"mini_project_restapi/router"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed load environment config")
	} else {
		fmt.Println("Success load environment config")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Connection failed:", err.Error())
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Connection failed:", err.Error())
	} else {
		fmt.Println("Connection success")
	}
	database.Initialize(DB)

	defer DB.Close()

	router.RunServer()
}

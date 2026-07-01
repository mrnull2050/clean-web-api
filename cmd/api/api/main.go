package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/mrnull2050/clean-web-api/docs"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mrnull2050/clean-web-api/internal/database"
	"github.com/mrnull2050/clean-web-api/internal/env"
)


// @title Go Gin REST API
// @version 1.0
// @descritpion  A rest API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your Bearer  token in this format  **Bearer Token**


type application struct {
	Port      int
	JWTSecret string
	models    database.Models
}

func main() {

	db, err := sql.Open("sqlite3", "/home/mr-null/Documents/Go-learning/REST-API/data.db")
	wd, _ := os.Getwd()
	fmt.Println("PWD:", wd)
	fmt.Println("DB path:", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	models := database.NewModel(db)
	app := &application{
		Port:      env.GetEnvInt("PORT", 8080),
		JWTSecret: env.GetEnvString("JWT-Secrest", "My-JWT-Secret-1234"),
		models:    models,
	}
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}

}

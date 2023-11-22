package main

import (
	"github/ersa97/med-record/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	APP_PORT string
	DB_URL   string
	db       *sqlx.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	APP_PORT = os.Getenv("APP_PORT")
	DB_URL = os.Getenv("DB_URL")

	db, err := sqlx.Connect("postgres", DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("ping database failed")
	} else {
		log.Println("database connection established")
	}
}

func main() {
	route := gin.Default()

	service := controllers.Controller{
		DB: db,
	}

	route.POST("/login", service.Login)
}

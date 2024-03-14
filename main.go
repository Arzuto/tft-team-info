package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"tft-team-info/controllers"
	"tft-team-info/database"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "aryarzt"
// 	dbname   = "tft-info"
// )

var (
	db  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	} else {
		fmt.Println("Successfully loaded.env file")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err = sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	database.DbMigrate(db)

	defer db.Close()
	router := gin.Default()
	router.GET("/origin", controllers.GetAllOrigin)
	router.POST("/origin", controllers.InsertOrigin)
	router.PATCH("/origin/:id", controllers.UpdateOrigin)
	router.DELETE("/origin/:id", controllers.DeleteOrigin)
	router.Run("localhost:8080")
}

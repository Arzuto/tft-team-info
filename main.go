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
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"))

	psqlInfo := fmt.Sprintf("host=%s  port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("PGHOST"),
	os.Getenv("PGPORT"),
	os.Getenv("PGUSER"),
	os.Getenv("PGPASSWORD"),
	os.Getenv("PGDATABASE"))

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
	router.POST("/origin", controllers.BasicAuthMiddleware(), controllers.InsertOrigin)
	router.PATCH("/origin/:id", controllers.BasicAuthMiddleware(), controllers.UpdateOrigin)
	router.DELETE("/origin/:id", controllers.BasicAuthMiddleware(), controllers.DeleteOrigin)

	router.GET("/class", controllers.GetAllClass)
	router.POST("/class", controllers.BasicAuthMiddleware(), controllers.InsertClass)
	router.PATCH("/class/:id", controllers.BasicAuthMiddleware(), controllers.UpdateClass)
	router.DELETE("/class/:id", controllers.BasicAuthMiddleware(), controllers.DeleteClass)

	router.GET("/item", controllers.GetAllItem)
	router.POST("/item", controllers.BasicAuthMiddleware(), controllers.InsertItem)
	router.PATCH("/item/:id", controllers.BasicAuthMiddleware(), controllers.UpdateItem)
	router.DELETE("/item/:id", controllers.BasicAuthMiddleware(), controllers.DeleteItem)

	router.GET("/character", controllers.GetAllCharacter)
	router.POST("/character", controllers.BasicAuthMiddleware(), controllers.InsertCharacter)
	router.PATCH("/character/:id", controllers.BasicAuthMiddleware(), controllers.UpdateCharacter)
	router.DELETE("/character/:id", controllers.BasicAuthMiddleware(), controllers.DeleteCharacter)

	router.GET("/recommendation", controllers.GetAllRecommendation)
	router.POST("/recommendation", controllers.BasicAuthMiddleware(), controllers.InsertRecommendation)
	router.PATCH("/recommendation/:id", controllers.BasicAuthMiddleware(), controllers.UpdateRecommendation)
	router.DELETE("/recommendation/:id", controllers.BasicAuthMiddleware(), controllers.DeleteRecommendation)

	// router.Run("localhost:8080")
	router.Run(":" + os.Getenv("PORT"))
}

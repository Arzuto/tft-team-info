package main

import (
	"database/sql"
	"fmt"
	"tft-team-info/database"

	_ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "aryarzt"
	dbname = "tft-info"
)

var(
	db *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password,dbname)

	db, err = sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}else{
		fmt.Println("Connected to database")
	}

	database.DbMigrate(db)
	
	defer db.Close()

}
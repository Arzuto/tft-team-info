package database

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/sql_migrations",
	}
	n, errs := migrate.Exec(dbParam,"postgres",migrations,migrate.Up)
	if errs != nil {
		panic(errs)
	}
	DbConnection = dbParam

	fmt.Println("Applied" ,n,(" migrations!"))
}
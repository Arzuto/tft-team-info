package database

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	dbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	dbConnection = dbParam
	migrations := &migrate.FileMigrationSource{
		Dir: "database/sql_migrations",
	}
	count, err := migrate.Exec(dbConnection, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", count)

	// count, err = migrate.Exec(dbConnection, "postgres", migrations, migrate.Down)
	// if err != nil {
	//     panic(err)
	// }
	// fmt.Printf("Table %d has been dropped.",count)
}
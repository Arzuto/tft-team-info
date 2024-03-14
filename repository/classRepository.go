package repository

import (
	"database/sql"
	"tft-team-info/structs"
)

func GetAllClass(db *sql.DB) ([]structs.Class, error) {
	sql := "SELECT * FROM class"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []structs.Class

	for rows.Next() {
		var classs structs.Class

		err := rows.Scan(&classs.ID, &classs.ClassName, &classs.Description)
		if err != nil {
			return nil, err
		}

		results = append(results, classs)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func InsertClass(db *sql.DB, class structs.Class) (err error) {
	sql := "INSERT INTO class (classname, description) VALUES ($1, $2)"

	errs := db.QueryRow(sql, class.ClassName, class.Description)

	return errs.Err()

}

func UpdateClass(db *sql.DB, class structs.Class) (err error) {
	sql := "UPDATE class SET classname = $1, description = $2 WHERE id = $3"

	errs := db.QueryRow(sql, class.ClassName, class.Description, class.ID)

	return errs.Err()

}

func DeleteClass(db *sql.DB, class structs.Class) (err error) {
	sql := "DELETE FROM class WHERE id = $1"

	errs := db.QueryRow(sql, class.ID)

	return errs.Err()

}

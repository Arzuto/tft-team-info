package repository

import (
	"database/sql"
	"tft-team-info/structs"
)

func GetAllOrigin(db *sql.DB) ([]structs.Origin, error) {
	sql := "SELECT * FROM origin"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []structs.Origin

	for rows.Next() {
		var origins structs.Origin

		err := rows.Scan(&origins.ID, &origins.OriginName, &origins.Description)
		if err != nil {
			return nil, err
		}

		results = append(results, origins)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func InsertOrigin(db *sql.DB, origin structs.Origin) (err error) {
	sql := "INSERT INTO origin (originname, description) VALUES ($1, $2)"

	errs := db.QueryRow(sql, origin.OriginName, origin.Description)

	return errs.Err()

}

func UpdateOrigin(db *sql.DB, origin structs.Origin) (err error) {
	sql := "UPDATE origin SET originname = $1, description = $2 WHERE id = $3"

	errs := db.QueryRow(sql, origin.OriginName, origin.Description, origin.ID)

	return errs.Err()

}

func DeleteOrigin(db *sql.DB, origin structs.Origin) (err error) {
	sql := "DELETE FROM origin WHERE id = $1"

	errs := db.QueryRow(sql, origin.ID)

	return errs.Err()

}

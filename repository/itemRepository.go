package repository

import (
	"database/sql"
	"tft-team-info/structs"
)

func GetAllItem(db *sql.DB) ([]structs.Item, error) {
	sql := "SELECT * FROM item"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []structs.Item

	for rows.Next() {
		var items structs.Item

		err := rows.Scan(&items.ID, &items.ItemName, &items.Description, &items.Stats)
		if err != nil {
			return nil, err
		}

		results = append(results, items)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func InsertItem(db *sql.DB, item structs.Item) (err error) {
	sql := "INSERT INTO item (itemname, description, stats) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, item.ItemName, item.Description, item.Stats)

	return errs.Err()

}

func UpdateItem(db *sql.DB, item structs.Item) (err error) {
	sql := "UPDATE item SET itemname = $1, description = $2, stats = $3 WHERE id = $4"

	errs := db.QueryRow(sql, item.ItemName, item.Description, item.Stats, item.ID)

	return errs.Err()

}

func DeleteItem(db *sql.DB, item structs.Item) (err error) {
	sql := "DELETE FROM item WHERE id = $1"

	errs := db.QueryRow(sql, item.ID)

	return errs.Err()

}

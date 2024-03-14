package repository

import (
	"database/sql"
	"tft-team-info/structs"
)

func GetAllCharacter(db *sql.DB) ([]structs.Character, error) {
	sql := "SELECT * FROM character"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []structs.Character

	for rows.Next() {
		var characters structs.Character

		err := rows.Scan(&characters.ID, &characters.CharacterName, &characters.Skill, &characters.OriginIDs, &characters.ClassIDs, &characters.ItemIDs)
		if err != nil {
			return nil, err
		}

		results = append(results, characters)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func InsertCharacter(db *sql.DB, character structs.Character) (err error) {
	sql := "INSERT INTO character (charactername, skill, origin_ids, class_ids, item_ids) VALUES ($1, $2, $3, $4, $5)"

	errs := db.QueryRow(sql, character.CharacterName, character.Skill, character.OriginIDs, character.ClassIDs, character.ItemIDs)

	return errs.Err()

}

func UpdateCharacter(db *sql.DB, character structs.Character) (err error) {
	sql := "UPDATE character SET charactername = $1, skill = $2, class_ids = $3, item_ids = $4 WHERE id = $5"

	errs := db.QueryRow(sql, character.CharacterName, character.Skill, character.ClassIDs, character.ItemIDs, character.ID)

	return errs.Err()

}

func DeleteCharacter(db *sql.DB, character structs.Character) (err error) {
	sql := "DELETE FROM character WHERE id = $1"

	errs := db.QueryRow(sql, character.ID)

	return errs.Err()

}

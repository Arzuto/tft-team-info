package repository

import (
	"database/sql"
	"fmt"
	"tft-team-info/structs"

	"github.com/lib/pq"
)

type IntSliceScanner struct {
	slice *[]int
}

func (scanner *IntSliceScanner) Scan(src interface{}) error {
	pqArray, ok := src.(pq.Int64Array)
	if !ok {
		return fmt.Errorf("src is not a pq.Int64Array")
	}

	ints := make([]int, len(pqArray))
	for i, v := range pqArray {
		ints[i] = int(v)
	}

	*scanner.slice = ints
	return nil
}

func GetAllCharacter(db *sql.DB) ([]structs.Character, error) {
	sql := `SELECT * from character`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []structs.Character

	for rows.Next() {
		var character structs.Character
		var classIDs pq.Int64Array
		var itemIDs pq.Int64Array

		err := rows.Scan(&character.ID, &character.CharacterName, &character.Skill, &character.OriginID, &classIDs, &itemIDs)
		if err != nil {
			return nil, err
		}

		character.ClassIDs = make([]int, len(classIDs))
		for i, id := range classIDs {
			character.ClassIDs[i] = int(id)
		}

		character.ItemIDs = make([]int, len(itemIDs))
		for i, id := range itemIDs {
			character.ItemIDs[i] = int(id)
		}

		results = append(results, character)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func InsertCharacter(db *sql.DB, character structs.Character) (err error) {
	sql := "INSERT INTO character (charactername, skill, originid, classids, itemids) VALUES ($1, $2, $3, $4, $5)"

	_, err = db.Exec(sql, character.CharacterName, character.Skill, character.OriginID, pq.Array(character.ClassIDs), pq.Array(character.ItemIDs))
	if err != nil {
		return err
	}

	return nil
}

func UpdateCharacter(db *sql.DB, character structs.Character) (err error) {
	sql := "UPDATE character SET charactername = $1, skill = $2, classids = $3, itemids = $4 WHERE id = $5"

	errs := db.QueryRow(sql, character.CharacterName, character.Skill, pq.Array(character.ClassIDs), pq.Array(character.ItemIDs), character.ID)

	return errs.Err()

}

func DeleteCharacter(db *sql.DB, character structs.Character) (err error) {
	sql := "DELETE FROM character WHERE id = $1"

	errs := db.QueryRow(sql, character.ID)

	return errs.Err()

}

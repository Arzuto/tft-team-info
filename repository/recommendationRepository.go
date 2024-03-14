package repository

import (
	"database/sql"
	"fmt"
	"tft-team-info/structs"

	"github.com/lib/pq"
)

type IntSliceScannerRecommendation struct {
	slice *[]int
}

func (scanner *IntSliceScannerRecommendation) Scan(src interface{}) error {
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

func GetAllRecommendation(db *sql.DB) ([]structs.Recommendation, error) {
	sql := `SELECT * from recommendation`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []structs.Recommendation

	for rows.Next() {
		var recommendation structs.Recommendation
		var originIDs pq.Int64Array
		var characterIDs pq.Int64Array
		var classIDs pq.Int64Array

		err := rows.Scan(&recommendation.ID, &recommendation.TeamName, &originIDs, &classIDs, &characterIDs, &recommendation.Tier, &recommendation.Difficulty)
		if err != nil {
			return nil, err
		}

		recommendation.OriginIDs = make([]int, len(originIDs))
		for i, id := range originIDs {
			recommendation.OriginIDs[i] = int(id)
		}
		recommendation.ClassIDs = make([]int, len(classIDs))
		for i, id := range classIDs {
			recommendation.ClassIDs[i] = int(id)
		}
		recommendation.CharacterIDs = make([]int, len(characterIDs))
		for i, id := range characterIDs {
			recommendation.CharacterIDs[i] = int(id)
		}

		results = append(results, recommendation)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func InsertRecommendation(db *sql.DB, recommendation structs.Recommendation) (err error) {
	sql := "INSERT INTO recommendation (teamname, originids, classids, characterids, tier, difficulty) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err = db.Exec(sql, recommendation.TeamName, pq.Array(recommendation.OriginIDs), pq.Array(recommendation.ClassIDs), pq.Array(recommendation.CharacterIDs), recommendation.Tier, recommendation.Difficulty)
	if err != nil {
		return err
	}

	return nil
}

func UpdateRecommendation(db *sql.DB, recommendation structs.Recommendation) (err error) {
	sql := "UPDATE recommendation SET teamname = $1, originids = $2, classids = $3, characterids = $4, tier = $5, difficulty = $6 WHERE id = $7"

	errs := db.QueryRow(sql, recommendation.TeamName, pq.Array(recommendation.OriginIDs), pq.Array(recommendation.ClassIDs), pq.Array(recommendation.CharacterIDs), recommendation.Tier, recommendation.Difficulty, recommendation.ID)

	return errs.Err()

}

func DeleteRecommendation(db *sql.DB, recommendation structs.Recommendation) (err error) {
	sql := "DELETE FROM recommendation WHERE id = $1"

	errs := db.QueryRow(sql, recommendation.ID)

	return errs.Err()

}

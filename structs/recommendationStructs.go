package structs

type Recommendation struct {
	ID           int64   `json:"id"`
	TeamName     string  `json:"team_name"`
	OriginIDs    []int64 `json:"origin_id"`
	ClassIDs     []int64 `json:"class_id"`
	CharacterIDs []int64 `json:"character_id"`
	Tier         string  `json:"tier"`
	Difficulty   string  `json:"difficulty"`
}

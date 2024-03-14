package structs

type Recommendation struct {
	ID           int64   `json:"id"`
	TeamName     string  `json:"team_name"`
	OriginIDs    []int `json:"origin_ids"`
	ClassIDs     []int `json:"class_ids"`
	CharacterIDs []int `json:"character_ids"`
	Tier         string  `json:"tier"`
	Difficulty   string  `json:"difficulty"`
}

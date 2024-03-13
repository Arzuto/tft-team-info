package structs

type Character struct {
	ID       int64   `json:"id"`
	CharacterName     string  `json:"character_name"`
	Skill    string  `json:"skill"`
	OriginID int64   `json:"origin_id"`
	ClassIDs  []int64   `json:"class_id"`
	ItemIDs  []int64 `json:"item_id"`
}

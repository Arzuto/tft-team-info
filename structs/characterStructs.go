package structs

type Character struct {
	ID       int64   `json:"id"`
	CharacterName     string  `json:"character_name"`
	Skill    string  `json:"skill"`
	OriginIDs int64   `json:"origin_ids"`
	ClassIDs  []int64   `json:"class_ids"`
	ItemIDs  []int64 `json:"item_ids"`
}

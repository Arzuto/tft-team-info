package structs

type Character struct {
	ID            int64  `json:"id"`
	CharacterName string `json:"character_name"`
	Skill         string `json:"skill"`
	OriginID      int    `json:"origin_id"`
	ClassIDs      []int  `json:"class_ids"`
	ItemIDs       []int  `json:"item_ids"`
}

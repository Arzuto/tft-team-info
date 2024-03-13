package structs

type Item struct {
	ID   int64  `json:"id"`
	ItemName string `json:"item_name"`
	Description string `json:"description"`
	Stats string `json:"stats"`
}
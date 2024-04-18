package dto

type ObjectPreviewDTO struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Tags  []string `json:"tags"`
	Score float32  `json:"score"`
}

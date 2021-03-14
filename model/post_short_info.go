package model

// PostShortInfo type for holding post items
type PostShortInfo struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	UpdatedDate string `json:"updated_date"`
}

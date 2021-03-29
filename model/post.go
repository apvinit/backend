package model

import "time"

// Post type for holding post items
type Post struct {
	ID           int64            `json:"id"`
	ShortLink    string           `json:"short_link,omitempty"`
	ImageLink    string           `json:"image_link,omitemtpy"`
	Type         string           `json:"type"`
	Title        string           `json:"title"`
	Name         string           `json:"name"`
	Info         string           `json:"info"`
	CreatedDate  string           `json:"created_date"`
	UpdatedDate  string           `json:"updated_date"`
	Organisation string           `json:"organisation"`
	TotalVacancy int              `json:"total_vacancy"`
	Dates        []ImportantDate  `json:"dates"`
	Links        []ImportantLink  `json:"links"`
	Fees         []ApplicationFee `json:"fees"`
	AgeLimits    []GeneralItem    `json:"age_limits"`
	AgeLimitAsOn string           `json:"age_limit_as_on"`
	Vacancies    []VacancyItem    `json:"vacancies"`
	Draft        bool             `json:"draft"`
	Trash        bool             `json:"trash"`
	CreatedAt    time.Time        `json:"created_at,omitempty"`
	UpdatedAt    time.Time        `json:"updated_at,omitempty"`
}

// ImportantDate is used for important date
type ImportantDate struct {
	ID    int    `json:"id"`
	Date  string `json:"date"`
	Title string `json:"title"`
}

// ImportantLink is used for important links
type ImportantLink struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// ApplicationFee is used for application fees
type ApplicationFee struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"amount"`
}

// GeneralItem is used for general items
type GeneralItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// VacancyItem is used for vacancy item
type VacancyItem struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Gen         string `json:"gen"`
	OBC         string `json:"obc"`
	BCA         string `json:"bca"`
	BCB         string `json:"bcb"`
	EWS         string `json:"ews"`
	SC          string `json:"sc"`
	ST          string `json:"st"`
	PH          string `json:"ph"`
	Total       string `json:"total"`
	AgeLimit    string `json:"age_limit"`
	Eligibility string `json:"eligibility"`
}

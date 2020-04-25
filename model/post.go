package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post type for holding post items
type Post struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type         string             `json:"type"`
	Title        string             `json:"title"`
	Name         string             `json:"name"`
	Info         string             `json:"info"`
	CreatedDate  string             `json:"createdDate"`
	UpdatedDate  string             `json:"updatedDate"`
	Organisation string             `json:"organisation"`
	TotalVacancy int                `json:"totalVacancy"`
	Dates        []ImportantDate    `json:"dates"`
	Links        []ImportantLink    `json:"links"`
	Fees         []ApplicationFee   `json:"fees"`
	AgeLimits    []GeneralItem      `json:"ageLimits"`
	AgeLimitAsOn string             `json:"ageLimitAsOn"`
	Vacancies    []VacancyItem      `json:"vacancies"`
	Draft        bool               `json:"draft"`
}

// ImportantDate is used for important date
type ImportantDate struct {
	Date  string `json:"date"`
	Title string `json:"title"`
}

// ImportantLink is used for important links
type ImportantLink struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// ApplicationFee is used for application fees
type ApplicationFee struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// GeneralItem is used for general items
type GeneralItem struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// VacancyItem is used for vacancy item
type VacancyItem struct {
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
	AgeLimit    string `json:"ageLimit"`
	Eligibility string `json:"eligibility"`
}

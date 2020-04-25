package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// PostShortInfo type for holding post items
type PostShortInfo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Type        string             `json:"type"`
	Title       string             `json:"title"`
	UpdatedDate string             `json:"updatedDate"`
}

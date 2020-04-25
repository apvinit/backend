package handler

import "go.mongodb.org/mongo-driver/mongo"

// Key for signing
const Key = "ThisIsASecret"

// Handler for managing mongo db access
type Handler struct {
	DB *mongo.Database
}

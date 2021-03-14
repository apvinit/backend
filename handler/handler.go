package handler

import (
	"database/sql"
)

// Key for signing
const Key = "ThisIsASecret"

// Handler for managing mongo db access
type Handler struct {
	DB *sql.DB
}

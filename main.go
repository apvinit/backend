package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sarkarijobadda/backend/conf"
	"github.com/sarkarijobadda/backend/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

const apikey1 = "wA5dZ8J1U4mt7X2LFRy9W8337Sda1eAotmSID8dYHHdUfer3"
const apikey2 = "diAENJWzWGdZmcS3M4/zOVZjSe0O9jhIdmVdG5uVXjasFlxr"
const apikey3 = "irabmvXNBCo3xf3bhRKagMwhOLbiLvlAlDkhqUIXC28ZTQNZ"

func main() {
	db, err := sql.Open("sqlite3", "file:jobadda.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	initDB(db)

	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} ${error}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == apikey1, nil
	}))

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Routes

	// Create update and delete post
	// Use auth middleware here
	e.POST("/posts", h.CreatePost)
	e.PUT("/posts/:id", h.UpdatePost)
	e.DELETE("/posts/:id", h.DeletePost)

	// Get one post
	e.GET("/posts/:id", h.FetchOnePost)
	// Get all the posts short info list (with type param)
	e.GET("/posts", h.GetPostShortInfo)
	// search endpoint (with q param)
	e.GET("/posts/search", h.SearchPost)

	// Start server
	e.HideBanner = true
	e.Logger.Fatal(e.Start(conf.ServerPort))

}

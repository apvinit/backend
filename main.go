package main

import (
	"backend/handler"
	"context"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseIP = "127.0.0.1"
var database = "test"

func main() {

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Mongodb database client options and connection
	clientOptions := options.Client().ApplyURI("mongodb://" + databaseIP + ":27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	println("Connected to database!")

	// Initialize handler
	h := &handler.Handler{DB: client.Database(database)}

	// Routes
	api := e.Group("/api")

	// Create update and delete post
	// Use auth middleware here
	api.POST("/post", h.CreatePost)
	api.PUT("/post/:id", h.UpdatePost)
	api.DELETE("/post/:id", h.DeletePost)

	// Get one post
	api.GET("/post/:id", h.FetchOnePost)
	// Get all the posts short info list (with type param)
	api.GET("/posts", h.GetPostShortInfo)
	// search endpoint (with q param)
	api.GET("/posts/search", h.SearchPost)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

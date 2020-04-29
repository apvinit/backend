package handler

import (
	"backend/conf"
	"backend/model"
	"backend/util"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreatePost method
func (h *Handler) CreatePost(c echo.Context) (err error) {
	p := &model.Post{
		ID: primitive.NewObjectID(),
	}

	if err = c.Bind(p); err != nil {
		return
	}

	if len(p.ImageLink) == 0 {
		p.ImageLink = conf.DefaultImageLink
	}

	shortLink, err := util.CreateDynamicLink(p)
	if err != nil {
		print(err)
	}
	p.ShortLink = shortLink

	_, err = h.DB.Collection("posts").InsertOne(context.TODO(), p)
	if err != nil {

		return
	}
	return c.JSON(http.StatusCreated, p)
}

// FetchPost method
func (h *Handler) FetchPost(c echo.Context) (err error) {
	posts := []*model.Post{}

	cur, err := h.DB.Collection("posts").Find(context.TODO(), bson.M{})

	if err != nil {
		return
	}
	for cur.Next(context.TODO()) {
		var p model.Post
		err = cur.Decode(&p)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, &p)
	}
	cur.Close(context.TODO())

	return c.JSON(http.StatusOK, posts)
}

// FetchOnePost method
func (h *Handler) FetchOnePost(c echo.Context) (err error) {
	_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": _id}

	var post model.Post

	err = h.DB.Collection("posts").FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, post)
}

// UpdatePost method
func (h *Handler) UpdatePost(c echo.Context) (err error) {
	_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": _id}

	p := &model.Post{}

	if err = c.Bind(p); err != nil {
		return
	}

	_, err = h.DB.Collection("posts").UpdateOne(context.TODO(), filter, bson.M{"$set": p})

	return c.NoContent(http.StatusOK)
}

// DeletePost method
func (h *Handler) DeletePost(c echo.Context) (err error) {

	_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": _id}

	_, err = h.DB.Collection("posts").DeleteOne(context.TODO(), filter)
	if err != nil {
		return
	}
	return c.NoContent(http.StatusOK)
}

// GetPostShortInfo  getting the short info for the posts
func (h *Handler) GetPostShortInfo(c echo.Context) (err error) {
	qp := c.QueryParam("type")

	var filter interface{}

	if len(qp) == 0 {
		filter = bson.M{}
	} else {
		filter = bson.M{"type": qp}
	}

	posts := []*model.PostShortInfo{}
	opt := options.Find().SetProjection(bson.M{}).SetLimit(100).SetSort(bson.M{"_id": -1})
	cur, err := h.DB.Collection("posts").Find(context.TODO(), filter, opt)

	if err != nil {
		return
	}
	for cur.Next(context.TODO()) {
		var p model.PostShortInfo
		err = cur.Decode(&p)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, &p)
	}
	cur.Close(context.TODO())
	return c.JSON(http.StatusOK, posts)
}

// SearchPost for searching post
func (h *Handler) SearchPost(c echo.Context) (err error) {
	q := c.QueryParam("q")
	println(q)
	posts := []*model.PostShortInfo{}
	opt := options.Find().SetProjection(bson.M{}).SetLimit(20).SetSort(bson.M{"_id": -1})
	cur, err := h.DB.Collection("posts").Find(context.TODO(), bson.M{"$text": bson.M{"$search": q}}, opt)
	if err != nil {
		log.Println(err)
	}
	for cur.Next(context.TODO()) {
		var p model.PostShortInfo
		err = cur.Decode(&p)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, &p)
	}
	cur.Close(context.TODO())
	return c.JSON(http.StatusOK, posts)
}

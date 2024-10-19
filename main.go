package main

import (
	// "database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BobadeKenny/BlogifyAPI/database"
	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.Default()
	database.ConnectDatabase()

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Docker! <3"})
	})
	g.POST("/posts", addPost)

	g.Run()
}

type Post struct {
	Title   string
	Slug    string
	Content string
}

func addPost(ctx *gin.Context) {
	body := Post{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
	}
	_, err = database.Db.Exec("insert into posts(title, slug, content) values ($1,$2,$3)", body.Title, body.Slug, body.Content)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Couldn't create post.")
	} else {
		ctx.JSON(http.StatusOK, "Post is successfully created.")
	}
}

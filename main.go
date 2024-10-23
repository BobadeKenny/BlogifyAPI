package main

import (
	// "database"
	"encoding/json"
	"fmt"
	"github.com/BobadeKenny/BlogifyAPI/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	Title     string
	Slug      string
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

func addPost(ctx *gin.Context) {
	body := Post{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Error 1 {err}")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
	}
	post := Post{Title: body.Title, Slug: body.Slug, Content: body.Content}
	result := database.Db.Create(&post)
	if result.Error != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, err)
	} else {
		ctx.JSON(http.StatusOK, "Post is successfully created.")
	}
}

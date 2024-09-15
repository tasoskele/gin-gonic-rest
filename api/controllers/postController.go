package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tasoskele/gin-gonic-rest/api/models"
	"github.com/tasoskele/gin-gonic-rest/api/utils"
)

func CreatePost(ctx *gin.Context) {
	//Request body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	ctx.Bind(&body)
	//Create POST request
	post := models.Post{Title: body.Title, Body: body.Body}
	result := utils.DB.Create(&post)

	// Result
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"message": "Error creating post",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Post created successfully",
		})
	}
}

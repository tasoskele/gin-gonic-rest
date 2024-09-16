package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tasoskele/gin-gonic-rest/api/models"
	"github.com/tasoskele/gin-gonic-rest/api/utils"
)

func GetProduct(ctx *gin.Context) {
	// Get single product
	var product models.Product
	id := ctx.Param("id")
	result := utils.DB.First(&product, id)

	// Result
	if result.Error != nil {
		ctx.JSON(404, gin.H{
			"message": "Product not found",
		})
	} else {
		ctx.JSON(200, product)
	}
}

func GetProducts(ctx *gin.Context) {
	// Get all products
	var products []models.Product
	result := utils.DB.Find(&products)

	// Result
	if result.Error != nil {
		ctx.JSON(404, gin.H{
			"message": "Products not found",
		})
	} else {
		ctx.JSON(200, products)
	}
}

func CreateProduct(ctx *gin.Context) {
	//Request body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	ctx.Bind(&body)
	//Create POST request
	product := models.Product{Title: body.Title, Body: body.Body}
	result := utils.DB.Create(&product)

	// Result
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"message": "Error creating product",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Product created successfully",
		})
	}
}

func UpdateProduct(ctx *gin.Context) {
	// Update product
	var product models.Product
	id := ctx.Param("id")
	result := utils.DB.First(&product, id)

	// Result
	if result.Error != nil {
		ctx.JSON(404, gin.H{
			"message": "Product not found",
		})
	} else {
		var body struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}
		ctx.Bind(&body)
		product.Title = body.Title
		product.Body = body.Body
		utils.DB.Save(&product)
		ctx.JSON(200, gin.H{
			"message": "Product updated successfully",
		})
	}
}

func DeleteProduct(ctx *gin.Context) {
	// Delete product
	var product models.Product
	id := ctx.Param("id")
	result := utils.DB.First(&product, id)

	// Result
	if result.Error != nil {
		ctx.JSON(404, gin.H{
			"message": "Product not found",
		})
	} else {
		utils.DB.Delete(&product)
		ctx.JSON(200, gin.H{
			"message": "Product deleted successfully",
		})
	}
}

package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tasoskele/gin-gonic-rest/api/models"
	"github.com/tasoskele/gin-gonic-rest/api/utils"
)

// GetProduct fetches a single product by ID
func GetProduct(ctx *gin.Context) {
	var product models.Product
	id := ctx.Param("id")
	result := utils.DB.First(&product, id)

	if result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// GetProducts fetches all products
func GetProducts(ctx *gin.Context) {
	var products []models.Product
	result := utils.DB.Find(&products)

	if result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving products",
		})
		return
	}

	if len(products) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No products found",
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func CreateProduct(ctx *gin.Context) {
	//Request body
	var body struct {
		Title string  `json:"title"`
		Body  string  `json:"body"`
		Price float64 `json:"price"`
	}
	ctx.Bind(&body)
	//Create POST request
	product := models.Product{Title: body.Title, Body: body.Body}
	result := utils.DB.Create(&product)

	// Result
	if result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(400, gin.H{
			"message": "Error creating product",
		})
	} else {
		ctx.JSON(201, gin.H{
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
		log.Println(result.Error)
		ctx.JSON(404, gin.H{
			"message": "Product not found",
		})
	} else {
		var body struct {
			Title string `json:"title"`
			Body  string `json:"body"`
			Price string `json:"price"`
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
		log.Println(result.Error)
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

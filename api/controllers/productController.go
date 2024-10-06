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

// GetProducts fetches all products / products index
func GetProducts(ctx *gin.Context) {
	var products []models.Product

	page, limit := utils.Paginate(ctx)
	result := utils.DB.Scopes(utils.PaginateScope(page, limit)).Find(&products)

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

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
		"page":     page,
		"limit":    limit,
	})
}

// CreateProduct creates a new product
func CreateProduct(ctx *gin.Context) {
	var body struct {
		Title string  `json:"title" binding:"required"`
		Body  string  `json:"body" binding:"required"`
		Price float64 `json:"price" binding:"required"`
	}
	ctx.Bind(&body)

	product := models.Product{
		Title: body.Title,
		Body:  body.Body,
		Price: body.Price,
	}
	result := utils.DB.Create(&product)

	if result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating product",
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Product created successfully",
			"product": product,
		})
	}
}

// UpdateProduct updates a product
func UpdateProduct(ctx *gin.Context) {
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

	var body struct {
		Title string
		Body  string
		Price float64
	}
	ctx.Bind(&body)

	product.Title = body.Title
	product.Body = body.Body
	product.Price = body.Price
	utils.DB.Save(&product)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": product,
	})
}

// DeleteProduct deletes a product
func DeleteProduct(ctx *gin.Context) {
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

	utils.DB.Unscoped().Delete(&product)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

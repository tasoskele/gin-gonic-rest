package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tasoskele/gin-gonic-rest/api/controllers"
	"github.com/tasoskele/gin-gonic-rest/api/models"
	"github.com/tasoskele/gin-gonic-rest/api/utils"
)

func init() {
	utils.ConnectDB()
	utils.DB.AutoMigrate(&models.Product{})
}

func main() {
	server := gin.Default()

	// Define your routes
	server.GET("/products", controllers.GetProducts)
	server.GET("/products/:id", controllers.GetProduct)
	server.POST("/products", controllers.CreateProduct)
	server.PUT("/products/:id", controllers.UpdateProduct)
	server.DELETE("/products/:id", controllers.DeleteProduct)

	// Get PORT from environment variables set by Docker, fallback to 8765
	port := os.Getenv("PORT")
	if port == "" {
		port = "8765"
	}

	// Start the server
	if err := server.Run(":" + port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

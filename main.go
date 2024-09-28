package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tasoskele/gin-gonic-rest/api/controllers"
	"github.com/tasoskele/gin-gonic-rest/api/utils"
)

func init() {
	utils.LoadEnv()
	utils.ConnectDB()
}

func main() {
	server := gin.Default()

	server.GET("/products", controllers.GetProducts)
	server.GET("/products/:id", controllers.GetProduct)
	server.POST("/products", controllers.CreateProduct)
	server.PUT("/products/:id", controllers.UpdateProduct)
	server.DELETE("/products/:id", controllers.DeleteProduct)

	server.Run()
}

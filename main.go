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

	server.POST("/products", controllers.CreateProduct)

	server.Run()
}

package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Test complete!",
		})
	})

	server.Run(":8081")
}

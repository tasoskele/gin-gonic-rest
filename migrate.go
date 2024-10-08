package main

import (
	"github.com/tasoskele/gin-gonic-rest/api/models"
	"github.com/tasoskele/gin-gonic-rest/api/utils"
)

func init() {
	utils.LoadEnv()
	utils.ConnectDB()
}

func RunMigrations() {
	utils.DB.AutoMigrate(&models.Product{})
}

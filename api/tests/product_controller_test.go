package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tasoskele/gin-gonic-rest/api/controllers"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.GetProduct)
	router.GET("/products", controllers.GetProducts)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
	return router
}

func TestProductsRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Paginate extracts pagination parameters from the query string
func Paginate(ctx *gin.Context) (int, int) {
	// Default values
	page := 1
	limit := 10

	// Parse 'page' query parameter
	if p := ctx.Query("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil {
			page = parsedPage
		}
	}

	// Parse 'limit' query parameter
	if l := ctx.Query("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil {
			limit = parsedLimit
		}
	}

	return page, limit
}

// PaginateScope applies GORM pagination
func PaginateScope(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

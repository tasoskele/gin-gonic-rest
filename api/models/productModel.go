package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	Price     float64 `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

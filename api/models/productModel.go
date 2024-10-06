package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title     string
	Body      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

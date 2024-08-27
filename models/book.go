package models

import "gorm.io/gorm"

// model for 'Book'

type Book struct {
	gorm.Model
	Title  string  `json:"title" gorm:"not null"`
	Author string  `json:"author"`
	Price  float64 `json:"price" gorm:"type:Decimal"`
}

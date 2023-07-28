package models

import "time"

type Product struct {
	ProductID   int       `gorm:"primary_key;auto_increment"`
	Name        string    `gorm:"varchar"`
	Description string    `gorm:"text"`
	Price       float64   `gorm:"decimal"`
	CreatedAt   time.Time `gorm:"datetime"`
	UpdatedAt   time.Time `gorm:"datetime"`
}

package models

import "time"

type Order struct {
	OrderID    int       `gorm:"primary_key;auto_increment"`
	TotalPrice float64   `gorm:"decimal"`
	CreatedAt  time.Time `gorm:"datetime"`
	UpdatedAt  time.Time `gorm:"datetime"`
	UserID     int       `gorm:"type:int"`
	OrderItems []OrderItem
}

type OrderItem struct {
	OrderItemID int       `gorm:"primary_key;auto_increment"`
	OrderID     int       `gorm:"type:int"`
	Quantity    int       `gorm:"int"`
	Price       float64   `gorm:"decimal"`
	CreatedAt   time.Time `gorm:"datetime"`
	UpdatedAt   time.Time `gorm:"datetime"`
	ProductID   int
	Product     Product
}

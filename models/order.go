package models

import "time"

type Order struct {
	OrderID    int       `gorm:"primary_key;auto_increment"`
	UserID     int       `gorm:"ForeignKey:UserID"`
	TotalPrice float64   `gorm:"decimal"`
	CreatedAt  time.Time `gorm:"datetime"`
	UpdatedAt  time.Time `gorm:"datetime"`
}

type OrderItem struct {
	OrderItemID int       `gorm:"primary_key;auto_increment"`
	OrderID     int       `gorm:"ForeignKey:OrderID"`
	ProductID   int       `gorm:"ForeignKey:ProductID"`
	Quantity    int       `gorm:"int"`
	CreatedAt   time.Time `gorm:"datetime"`
	UpdatedAt   time.Time `gorm:"datetime"`
}

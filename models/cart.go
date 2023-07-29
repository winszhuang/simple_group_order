package models

import "time"

type Cart struct {
	CartID    int       `gorm:"primary_key;auto_increment"`
	UserID    int       `gorm:"type:int"`
	CreatedAt time.Time `gorm:"datetime"`
	UpdatedAt time.Time `gorm:"datetime"`
	CartItems []CartItem
}

type CartItem struct {
	CartItemID int       `gorm:"primary_key;auto_increment"`
	CartID     int       `gorm:"type:int"`
	Quantity   int       `gorm:"int"`
	CreatedAt  time.Time `gorm:"datetime"`
	UpdatedAt  time.Time `gorm:"datetime"`
	ProductID  int
	Product    Product
}

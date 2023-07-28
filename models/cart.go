package models

import "time"

type Cart struct {
	CartID    int       `gorm:"primary_key;auto_increment"`
	UserID    int       `gorm:"ForeignKey:UserID"`
	CreatedAt time.Time `gorm:"datetime"`
	UpdatedAt time.Time `gorm:"datetime"`
}

type CartItem struct {
	CartItemID int       `gorm:"primary_key;auto_increment"`
	CartID     int       `gorm:"ForeignKey:CartID"`
	ProductID  int       `gorm:"ForeignKey:ProductID"`
	Quantity   int       `gorm:"int"`
	CreatedAt  time.Time `gorm:"datetime"`
	UpdatedAt  time.Time `gorm:"datetime"`
}

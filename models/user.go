package models

import "time"

type User struct {
	UserID       int       `gorm:"primary_key;auto_increment"`
	Username     string    `gorm:"varchar"`
	Password     string    `gorm:"varchar"`
	Email        string    `gorm:"varchar"`
	Balance      float64   `gorm:"decimal"` // 新增的餘額欄位
	CreatedAt    time.Time `gorm:"datetime"`
	Carts        []Cart
	Orders       []Order
	Transactions []Transaction
}

package models

import "time"

type Transaction struct {
	TransactionID int       `gorm:"primary_key;auto_increment"`
	UserID        int       `gorm:"type:int"`
	Amount        float64   `gorm:"decimal"` // 變動的金額，正數表示增加，負數表示減少
	CreatedAt     time.Time `gorm:"datetime"`
}

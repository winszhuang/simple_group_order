package models

import "time"

type Product struct {
	ProductID   int             `gorm:"primary_key;auto_increment"`
	Name        string          `gorm:"varchar"`
	Description string          `gorm:"text"`
	Price       float64         `gorm:"decimal"`
	Image       string          `gorm:"text"`
	CreatedAt   time.Time       `gorm:"datetime"`
	BrandID     int             `gorm:"type:int"`
	Brand       ProductBrand    `gorm:"foreignKey:BrandID"`
	CategoryID  int             `gorm:"type:int"`
	Category    ProductCategory `gorm:"foreignKey:CategoryID"`
}

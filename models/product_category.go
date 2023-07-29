package models

type ProductCategory struct {
	ProductCategoryID int    `gorm:"primary_key;auto_increment"`
	Name              string `gorm:"varchar"`
	Description       string `gorm:"text"`
}

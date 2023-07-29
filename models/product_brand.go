package models

type ProductBrand struct {
	ProductBrandID int    `gorm:"primary_key;auto_increment"`
	Name           string `gorm:"varchar"`
	Description    string `gorm:"text"`
}

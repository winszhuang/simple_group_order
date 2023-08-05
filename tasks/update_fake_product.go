package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"os"
	"simple_group_order/db"
	"simple_group_order/models"
)

type Data struct {
	Products []Product `json:"products"`
	Total    int64     `json:"total"`
	Skip     int64     `json:"skip"`
	Limit    int64     `json:"limit"`
}

type Product struct {
	ID                 int64    `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int64    `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int64    `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := db.NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	if err = database.Migrate(); err != nil {
		panic(err)
	}
	defer database.Close()

	err = updateProductsToDB(database.GetDB())
	if err != nil {
		panic(err)
	}
}

func updateProductsToDB(db *gorm.DB) error {
	source, err := os.ReadFile("products.json")
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(source, &data); err != nil {
		return err
	}

	for _, product := range data.Products {
		brand := getOrCreateBrand(db, product.Brand)
		category := getOrCreateCategory(db, product.Category)

		p := models.Product{
			Name:        product.Title,
			Description: product.Description,
			Price:       float64(product.Price),
			Image:       product.Images[0],
			BrandID:     brand.ProductBrandID,
			CategoryID:  category.ProductCategoryID,
		}

		result := db.Create(&p)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func getOrCreateBrand(db *gorm.DB, brandName string) *models.ProductBrand {
	var brand models.ProductBrand
	db.Where("name = ?", brandName).FirstOrCreate(&brand, models.ProductBrand{
		Name:        brandName,
		Description: "",
	})
	return &brand
}

func getOrCreateCategory(db *gorm.DB, categoryName string) *models.ProductCategory {
	var brand models.ProductCategory
	db.Where("name = ?", categoryName).FirstOrCreate(&brand, models.ProductCategory{
		Name:        categoryName,
		Description: "",
	})
	return &brand
}

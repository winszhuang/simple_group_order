package main

import (
	"encoding/json"
	"fmt"
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
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := db.NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	if err = database.Migrate(); err != nil {
		panic(err)
	}

	err = updateProductsToDB(database.GetDB())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("更新假資料成功")
	}
}

func updateProductsToDB(db *gorm.DB) error {
	source, err := os.ReadFile("tasks/products.json")
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(source, &data); err != nil {
		return err
	}

	for _, product := range data.Products {
		brand, err := getOrCreateBrand(db, product.Brand)
		if err != nil {
			return err
		}
		category, err := getOrCreateCategory(db, product.Category)
		if err != nil {
			return err
		}

		p := models.Product{
			Name:        product.Title,
			Description: product.Description,
			Price:       float64(product.Price),
			Image:       product.Images[0],
			BrandID:     brand.ProductBrandID,
			CategoryID:  category.ProductCategoryID,
		}

		err = db.Create(&p).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func getOrCreateBrand(db *gorm.DB, brandName string) (*models.ProductBrand, error) {
	var brand models.ProductBrand
	if err := db.Where("name = ?", brandName).FirstOrCreate(&brand, models.ProductBrand{
		Name:        brandName,
		Description: "",
	}).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}

func getOrCreateCategory(db *gorm.DB, categoryName string) (*models.ProductCategory, error) {
	var brand models.ProductCategory
	if err := db.Where("name = ?", categoryName).FirstOrCreate(&brand, models.ProductCategory{
		Name:        categoryName,
		Description: "",
	}).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}

package db

import (
	"simple_group_order/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &Database{db: database}, nil
}

func (d *Database) Migrate() error {
	return d.db.AutoMigrate(
		&models.User{},
		&models.Transaction{},
		&models.Cart{},
		&models.CartItem{},
		&models.Product{},
		&models.ProductCategory{},
		&models.ProductBrand{},
		&models.Order{},
		&models.OrderItem{},
	)
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

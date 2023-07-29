package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type CartController struct {
	db *gorm.DB
}

func NewCartController(db *gorm.DB) *CartController {
	return &CartController{db: db}
}

func (cc *CartController) CreateCart(c *gin.Context) {
	//cart := &models.Cart{}
	//cc.db.Create(&cart)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (cc *CartController) GetCartData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (cc *CartController) AddToCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (cc *CartController) RemoveCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (cc *CartController) RemoveCartItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (cc *CartController) UpdateCartItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (cc *CartController) EmptyCartData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (cc *CartController) Checkout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

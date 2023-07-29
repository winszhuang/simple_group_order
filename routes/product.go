package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"simple_group_order/controllers"
	"simple_group_order/middlewares"
)

func AddProductRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	productController := controllers.NewProductController(db)
	product := rg.Group("/product", middlewares.Authorize())
	product.GET("/", productController.List)
	product.GET("/:id", productController.Get)
}

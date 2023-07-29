package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"simple_group_order/controllers"
	"simple_group_order/middlewares"
)

func AddCartRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	cartController := controllers.NewCartController(db)
	cart := rg.Group("/cart", middlewares.Authorize())
	cart.POST("/create", cartController.CreateCart)
	cart.GET("/:cartId", cartController.GetCartData)
	cart.POST("/:cartId", cartController.AddToCart)
	cart.DELETE("/:cartId", cartController.RemoveCart)
	cart.DELETE("/:cartId/item/:itemId", cartController.RemoveCartItem)
	cart.PATCH("/:cartId/item/:itemId", cartController.UpdateCartItem)
	cart.DELETE("/:cartId/empty", cartController.EmptyCartData)
	cart.POST("/:cartId/checkout", cartController.Checkout)
}

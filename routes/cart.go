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

	cartItem := cart.Group("/:cartId", middlewares.CheckUserCart(db))
	cartItem.GET("", cartController.GetCartData)
	cartItem.POST("", cartController.AddToCart)
	cartItem.DELETE("", cartController.RemoveCart)
	cartItem.DELETE("/item/:itemId", cartController.RemoveCartItem)
	cartItem.PATCH("/item/:itemId", cartController.UpdateCartItem)
	cartItem.DELETE("/empty", cartController.EmptyCartData)
	cartItem.POST("/checkout", cartController.Checkout)
}

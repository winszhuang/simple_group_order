package routes

import (
	"simple_group_order/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAuthRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	authController := controllers.NewAuthController(db)
	auth := rg.Group("/auth")
	auth.POST("/signup", authController.SignUp)
	auth.POST("/signin", authController.SignIn)
	auth.POST("/signout", authController.SignOut)
}

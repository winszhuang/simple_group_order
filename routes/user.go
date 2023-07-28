package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"simple_group_order/controllers"
	"simple_group_order/middlewares"
)

func AddUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userController := controllers.NewUserController(db)
	user := rg.Group("/user", middlewares.Authorize())
	user.GET("/me", userController.GetUserInfo)
}

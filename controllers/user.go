package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"simple_group_order/models"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")

	// get user
	user := &models.User{}
	uc.db.Where("user_id = ?", userId).First(&user)

	// check user is valid
	invalidUser := user.UserID == 0
	if invalidUser {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "找不到該使用者",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email":  user.Email,
		"userId": user.UserID,
		"name":   user.Username,
	})
}

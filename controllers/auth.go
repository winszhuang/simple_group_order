package controllers

import (
	"fmt"
	"net/http"
	"simple_group_order/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createUserRequest struct {
	UserName string `json:"username" binding:"min=2,max=10"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"min=4,max=10"`
}

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		db: db,
	}
}

func (ac *AuthController) SignUp(c *gin.Context) {
	request := &createUserRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "參數錯誤",
			"error":   err.Error(),
		})
		return
	}

	user := models.User{}
	ac.db.Where("username = ?", request.UserName).First(&user)

	existedUser := user.UserID != 0
	if existedUser {
		message := fmt.Sprintf("使用者暱稱 {%s} 已被取過，請重新", request.UserName)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	ac.db.Where("email = ?", request.Email).First(&user)
	existedUser = user.UserID != 0
	if existedUser {
		message := fmt.Sprintf("使用者email {%s} 已被註冊，請更換email", request.Email)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	user = models.User{
		Username:  request.UserName,
		Password:  request.Password,
		Email:     request.Email,
		CreatedAt: time.Now(),
	}
	result := ac.db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "註冊成功",
		"userId":  user.UserID,
	})
}

func (ac *AuthController) SignIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "test",
	})
}

func (ac *AuthController) SignOut(c *gin.Context) {
	return
}

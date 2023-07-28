package controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"simple_group_order/models"
	"simple_group_order/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignUpRequest struct {
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
	request := &SignUpRequest{}
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
		message := fmt.Sprintf("使用者暱稱 {%s} 已被取過，請變更名稱", request.UserName)
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

	// hash password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	user = models.User{
		Username:  request.UserName,
		Password:  string(hashPassword),
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

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"min=4,max=10"`
}

func (ac *AuthController) SignIn(c *gin.Context) {
	request := &SignInRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "參數錯誤",
			"error":   err.Error(),
		})
		return
	}

	// find user
	user := models.User{}
	ac.db.Where("email = ?", request.Email).First(&user)

	// check user existed
	invalidUser := user.UserID == 0
	if invalidUser {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "無此使用者，請先註冊",
		})
		return
	}

	// validate password
	hashPassword := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密碼錯誤",
			"error":   err.Error(),
		})
		return
	}

	// generate jwt
	jwtStr, err := utils.CreateJWT(map[string]string{
		"userId": strconv.Itoa(user.UserID),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統內部錯誤",
		})
		println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "登入成功",
		"token":   jwtStr,
	})
}

func (ac *AuthController) SignOut(c *gin.Context) {
	return
}

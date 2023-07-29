package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"simple_group_order/constants"
	"simple_group_order/models"
	"simple_group_order/utils"
	"strconv"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{db: db}
}

func (pc *ProductController) List(c *gin.Context) {
	var pageIndex, pageSize int
	pageIndexStr := c.Query("pageIndex")
	pageSizeStr := c.Query("pageSize")
	if !utils.IsPositiveInteger(pageIndexStr) && pageIndexStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "錯誤的pageIndex參數，請輸入大於0的數字",
		})
		return
	}
	if !utils.IsPositiveInteger(pageSizeStr) && pageSizeStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "錯誤的pageSize參數，請輸入大於0的數字",
		})
		return
	}

	if pageIndexStr == "" {
		pageIndex = constants.DefaultPageIndex
	} else {
		pageIndex, _ = strconv.Atoi(pageIndexStr)
	}
	if pageSizeStr == "" {
		pageSize = constants.DefaultPageSize
	} else {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	var productList []models.Product
	offset := (pageIndex - 1) * pageSize
	pc.db.Limit(pageSize).Offset(offset).Find(&productList)

	c.JSON(http.StatusOK, gin.H{
		"data":      productList,
		"pageIndex": pageIndex,
		"pageSize":  pageSize,
	})
}

func (pc *ProductController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "錯誤的id",
			"error":   err.Error(),
		})
		return
	}

	var product models.Product
	result := pc.db.First(&product, id)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"data":    "",
			"message": "查無此商品",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}

}

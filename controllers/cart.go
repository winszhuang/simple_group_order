package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"simple_group_order/models"
	"simple_group_order/types"
	"simple_group_order/utils"
	"strconv"
)

type CartController struct {
	db *gorm.DB
}

func NewCartController(db *gorm.DB) *CartController {
	return &CartController{db: db}
}

func (cc *CartController) CreateCart(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "錯誤的使用者",
		})
		return
	}

	cart := &models.Cart{UserID: userId.(int)}
	if err := cc.db.Create(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "創建購物車失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cartId": cart.CartID,
	})
}

func (cc *CartController) GetCartData(c *gin.Context) {
	cartData, _ := c.Get("cartData")

	c.JSON(http.StatusOK, gin.H{
		"data": cartData,
	})
}

type AddToCartRequest struct {
	Quantity  int `json:"quantity" binding:"min=1"`
	ProductID int `json:"productID" binding:"min=1"`
}

func (cc *CartController) AddToCart(c *gin.Context) {
	data, _ := c.Get("cartData")
	cartData := data.(types.CartData)

	//取得body內需要加入購物車的資料
	request := &AddToCartRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "參數錯誤",
			"error":   err.Error(),
		})
		return
	}

	cartItem := models.CartItem{
		CartID:    cartData.CartID,
		Quantity:  request.Quantity,
		ProductID: request.ProductID,
	}

	//創建購物車前，判斷同商品是否已在購物車內，是的話增加商品數量
	var existedCartItem *types.CartItemBaseData
	for _, item := range cartData.Items {
		if item.ProductID == request.ProductID {
			existedCartItem = &item
			break
		}
	}
	if existedCartItem != nil {
		cartItem.CartItemID = existedCartItem.CartItemID
		newQuantity := existedCartItem.Quantity + request.Quantity
		err := cc.db.Model(&cartItem).Update("quantity", newQuantity).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "更新購物車品項資訊失敗",
				"error":   err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "更新商品數量成功!!",
				"data":    cartItem,
			})
		}
		return
	}

	//創建購物車item，並加入資料庫
	if err := cc.db.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "創建購物車品項失敗",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "創建成功!!",
			"data":    cartItem,
		})
	}
}

func (cc *CartController) RemoveCart(c *gin.Context) {
	data, _ := c.Get("cartData")
	cartData := data.(types.CartData)

	tx := cc.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "刪除購物車錯誤",
				"error":   r.(error).Error(),
			})
		}
	}()

	if err := tx.Error; err != nil {
		panic(err)
	}

	if err := tx.Where("cart_id = ?", cartData.CartID).Delete(&models.CartItem{}).Error; err != nil {
		panic(err)
	}

	if err := tx.Delete(&models.Cart{}, cartData.CartID).Error; err != nil {
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "刪除購物車成功",
	})
}

func (cc *CartController) RemoveCartItem(c *gin.Context) {
	itemIDStr := c.Param("itemId")
	if !utils.IsPositiveInteger(itemIDStr) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "item錯誤!! itemID需要為正整數",
		})
		return
	}

	itemID, _ := strconv.Atoi(itemIDStr)
	cartItem := models.CartItem{}
	if err := cc.db.First(&cartItem, itemID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查無此購物品項",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.db.Delete(&cartItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "刪除品項失敗!!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "刪除品項成功",
		})
	}
}

type UpdateCartItemRequest struct {
	Quantity int
}

func (cc *CartController) UpdateCartItem(c *gin.Context) {
	itemIDStr := c.Param("itemId")
	if !utils.IsPositiveInteger(itemIDStr) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "item錯誤!! itemID需要為正整數",
		})
		return
	}

	itemID, _ := strconv.Atoi(itemIDStr)
	request := &UpdateCartItemRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "參數錯誤",
			"error":   err.Error(),
		})
		return
	}

	if request.Quantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "數量不得小於1",
		})
		return
	}

	cartItem := models.CartItem{}
	if err := cc.db.First(&cartItem, itemID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查無此購物品項",
			"error":   err.Error(),
		})
		return
	}

	err := cc.db.Model(&cartItem).Update("quantity", request.Quantity).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "更新商品數量失敗",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "更新數量成功",
			"data":    cartItem,
		})
	}
}

func (cc *CartController) EmptyCartData(c *gin.Context) {
	data, _ := c.Get("cartData")
	cartData := data.(types.CartData)

	tx := cc.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "清空購物車內商品失敗",
				"error":   r.(error).Error(),
			})
		}
	}()

	if err := tx.Where("cart_id = ?", cartData.CartID).Delete(&models.CartItem{}).Error; err != nil {
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "清空購物車內商品成功",
	})
}

func (cc *CartController) Checkout(c *gin.Context) {
	data, _ := c.Get("cartData")
	cartData := data.(types.CartData)

	tx := cc.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "結帳失敗",
				"error":   r.(error).Error(),
			})
		}
	}()

	if err := tx.Error; err != nil {
		panic(err)
	}

	order := models.Order{UserID: cartData.UserID}
	if err := tx.Create(&order).Error; err != nil {
		panic(err)
	}

	var totalPrice float64
	var orderItems []models.OrderItem
	for _, item := range cartData.Items {
		product := models.Product{}
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			panic(err)
		}

		orderItem := models.OrderItem{
			OrderID:   order.OrderID,
			Quantity:  item.Quantity,
			Price:     product.Price,
			ProductID: product.ProductID,
		}
		if err := tx.Create(&orderItem).Error; err != nil {
			panic(err)
		}

		orderItems = append(orderItems, orderItem)

		// 更新總價
		totalPrice = totalPrice + product.Price*float64(item.Quantity)
	}

	// 更新order的總價錢
	if err := tx.Model(&order).Update("total_price", totalPrice).Error; err != nil {
		panic(err)
	}

	// 更新使用者餘額
	user := models.User{}
	if err := tx.First(&user, cartData.UserID).Error; err != nil {
		panic(err)
	}

	newBalance := user.Balance - totalPrice
	if newBalance < 0 {
		panic(errors.New("餘額不足!!"))
	}

	if err := tx.Model(user).Update("balance", newBalance).Error; err != nil {
		panic(err)
	}

	// 刪除不需要的購物車和購物車品項
	if err := tx.Where("cart_id = ?", cartData.CartID).Delete(&models.CartItem{}).Error; err != nil {
		panic(err)
	}
	if err := tx.Delete(&models.Cart{}, cartData.CartID).Error; err != nil {
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "結帳成功!!",
		"order":      order,
		"orderItems": orderItems,
	})
}

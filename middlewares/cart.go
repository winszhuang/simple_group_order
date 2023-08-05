package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"simple_group_order/models"
	"simple_group_order/types"
	"simple_group_order/utils"
	"strconv"
)

func CheckUserCart(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := c.Get("userId")

		//先判斷該購物車id是否為該使用者有的購物車id，並記錄該cartId
		cartIdStr := c.Param("cartId")
		if !utils.IsPositiveInteger(cartIdStr) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "錯誤的購物車代號!格式為正整數",
			})
			return
		}

		//確認是否有該購物車
		cartId, _ := strconv.Atoi(cartIdStr)
		var cart models.Cart
		result := db.First(&cart, cartId)
		if result.Error == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "查無此購物車代號",
				"error":   result.Error.Error(),
			})
			return
		}

		//判斷該購物車是否為該使用者的
		if cart.UserID != userId {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "使用者無此購物車",
			})
			return
		}

		//取出購物車內所有項目並儲存進context
		var cartItems []types.CartItemBaseData
		db.Table("carts").Select("cart_item_id, quantity, product_id").Joins("JOIN cart_items ON cart_items.cart_id = carts.cart_id").Scan(&cartItems)
		c.Set("cartData", types.CartData{
			CartID: cart.CartID,
			UserID: cart.UserID,
			Items:  cartItems,
		})
	}
}

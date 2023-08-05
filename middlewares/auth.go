package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_group_order/utils"
	"strconv"
	"strings"
)

const (
	authPrefix = "Bearer"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "沒有授權，請提供授權")
			return
		}

		if !strings.HasPrefix(bearerToken, authPrefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "錯誤的授權格式")
			return
		}

		strArr := strings.Split(bearerToken, " ")
		if len(strArr) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "錯誤的授權格式")
			return
		}

		tokenStr := strArr[1]
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "授權為空!!")
			return
		}

		claims, err := utils.VerifyJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "錯誤的授權!!")
			return
		}

		userIdStr := claims["userId"].(string)
		userId, _ := strconv.Atoi(userIdStr)
		c.Set("userId", userId)
	}
}

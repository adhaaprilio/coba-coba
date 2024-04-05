package middleware

import (
	"backend/errorHandler"
	"backend/helper"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorHandler.HandleError(c, &errorHandler.Error401{Message: "Unathorized"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)

		if err != nil {
			errorHandler.HandleError(c, &errorHandler.Error401{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userId)
		c.Next()
	}
}

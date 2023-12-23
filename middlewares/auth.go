package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/IvanOrsh/go-rest-event-booking/utils"
)

func Authenticate(c *gin.Context) {
	// check whether the token exists
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
				"error":   "no token provided",
			})
		return
	}

	// check whether the token is valid
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
				"error":   err.Error(),
			})
		return
	}

	c.Set("userId", userId)

	c.Next()
}

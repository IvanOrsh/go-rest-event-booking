package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/IvanOrsh/go-rest-event-booking/models"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse request data",
				"error":   err.Error(),
			})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not create user",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "user created",
			"data":    user,
		},
	)
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse request data",
				"error":   err.Error(),
			})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "invalid credentials",
				"error":   err.Error(),
			})
		return
	}

	// TODO: return JWT
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "login successful",
			"data":    user,
		},
	)

}

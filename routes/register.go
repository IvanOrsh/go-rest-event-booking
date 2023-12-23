package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/IvanOrsh/go-rest-event-booking/models"
)

func register(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse id",
				"error":   err.Error(),
			})
		return
	}

	// check if event exists
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not retrieve event",
				"error":   err.Error(),
			})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not register user for event",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "user registered for event",
		},
	)
}

func cancel(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse id",
				"error":   err.Error(),
			})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.Unregister(userId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not unregister user for event",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "user unregistered for event",
		},
	)

}

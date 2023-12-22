package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/IvanOrsh/go-rest-event-booking/models"
	"github.com/IvanOrsh/go-rest-event-booking/utils"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not retrieve events",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "events retrieved",
			"data":    events,
		})
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse id",
				"error":   err.Error(),
			})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not retrieve event",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "event retrieved",
			"data":    event,
		},
	)
}

func createEvent(c *gin.Context) {
	// check whether the token exists
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
				"error":   "no token provided",
			})
		return
	}

	// check whether the token is valid
	err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
				"error":   err.Error(),
			})
		return
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse request data",
				"error":   err.Error(),
			})
		return
	}
	event.UserID = 1

	err = event.Save()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not create event",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "event created",
			"data":    event,
		},
	)
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse id",
				"error":   err.Error(),
			})
		return
	}

	_, err = models.GetEventByID(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not retrieve event",
				"error":   err.Error(),
			})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse request data",
				"error":   err.Error(),
			})
		return
	}

	updatedEvent.ID = id
	err = models.Update(updatedEvent)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not update event",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{

			"message": "event updated",
		},
	)
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse id",
				"error":   err.Error(),
			})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not retrieve event",
				"error":   err.Error(),
			})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not delete event",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "event deleted",
		},
	)
}

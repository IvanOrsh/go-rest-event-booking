package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/IvanOrsh/go-rest-event-booking/db"
	"github.com/IvanOrsh/go-rest-event-booking/models"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

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

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "could not parse request data",
				"error3":  err.Error(),
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

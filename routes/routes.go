package routes

import (
	"github.com/IvanOrsh/go-rest-event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", register)
	authenticated.DELETE("/events/:id/cancel", cancel)

	server.POST("/signup", signup)
	server.POST("/login", login)
}

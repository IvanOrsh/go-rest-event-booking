package main

import (
	"github.com/gin-gonic/gin"

	"github.com/IvanOrsh/go-rest-event-booking/db"
	"github.com/IvanOrsh/go-rest-event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

package routes

import (
	"github.com/gin-gonic/gin"
	"www.example.com/booking/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/event", getEvents)
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authentication)
	authenticated.POST("/event", createEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)

	server.POST("/signup", singup)
	server.POST("/login", login)
}

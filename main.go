package main

import (
	"github.com/gin-gonic/gin"
	"www.example.com/booking/db"
	"www.example.com/booking/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}

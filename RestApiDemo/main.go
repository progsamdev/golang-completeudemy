package main

import (
	"log"
	"restapidemo/db"
	"restapidemo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

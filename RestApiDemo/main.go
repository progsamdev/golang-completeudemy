package main

import (
	"log"
	"net/http"
	"restapidemo/Models"
	"restapidemo/db"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := Models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var newEvent Models.Event

	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err) //log error but do not stop the program
		return
	}

	newEvent.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

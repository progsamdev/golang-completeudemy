package main

import (
	"log"
	"net/http"
	"restapidemo/Models"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func main() {
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

	newEvent.UUID = uuid.Must(uuid.NewV4())
	newEvent.UserID = uuid.Must(uuid.NewV4())
	newEvent.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

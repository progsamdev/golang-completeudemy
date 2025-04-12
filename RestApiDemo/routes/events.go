package routes

import (
	"log"
	"net/http"
	"restapidemo/Models"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func getEvents(c *gin.Context) {
	events, err := Models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var newEvent Models.Event

	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err) //log error but do not stop the program
		return
	}

	err := newEvent.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

func getEventById(c *gin.Context) {
	id := c.Param("id") //get id from url
	uuid, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	event, err := Models.GetEventById(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

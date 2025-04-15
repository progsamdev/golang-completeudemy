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

	userID := c.GetString("userID")

	var newEvent Models.Event

	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	newEvent.UserID = uuid.Must(uuid.FromString(userID))
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

func updateEvent(c *gin.Context) {
	id := c.Param("id")
	eventId, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var updatedEvent *Models.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	userId := c.GetString("userID")
	parsedUserId, err := uuid.FromString(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	updatedEvent, err = Models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if updatedEvent.UserID != parsedUserId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(c *gin.Context) {
	userId := c.GetString("userID")
	parsedUserId, err := uuid.FromString(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	id := c.Param("id")
	eventId, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	event, err := Models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != parsedUserId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

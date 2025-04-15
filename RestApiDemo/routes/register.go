package routes

import (
	"net/http"
	"restapidemo/Models"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func registerForEvent(c *gin.Context) {
	eventId := c.Param("id")
	userId := c.GetString("userID")

	parsedEventId, err := uuid.FromString(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event id"})
		return
	}

	event, err := Models.GetEventById(parsedEventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch the event."})
		return
	}

	err = event.Register(uuid.Must(uuid.FromString(userId)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for the event."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered for the event successfully."})
}

func unregisterForEvent(c *gin.Context) {
	eventId := c.Param("id")
	userId := c.GetString("userID")

	parsedEventId, err := uuid.FromString(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event id"})
		return
	}

	event, err := Models.GetEventById(parsedEventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch the event."})
		return
	}

	err = event.Unregister(uuid.Must(uuid.FromString(userId)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not unregister for the event."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Unregistered for the event successfully."})
}

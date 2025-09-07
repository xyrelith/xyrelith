package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"xyrelith/api/models"
	"xyrelith/api/postgres"
)

func CreateEvent(c *gin.Context) {
	// bind JSON data from POST request to variable
	var newEvent models.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event := models.Event{
		Title:     newEvent.Title,
		StartDate: newEvent.StartDate,
		EndDate:   newEvent.EndDate,
		Priority:  newEvent.Priority,
	}

	// Add event to db, on error return err
	if result := x_postgres.DB.Create(&event); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

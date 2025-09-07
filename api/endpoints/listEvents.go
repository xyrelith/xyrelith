package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"xyrelith/api/models"
	"xyrelith/api/postgres"
)

func ListEvents(c *gin.Context) {
	var events []models.Event

	dbResult := x_postgres.DB.Find(&events)
	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

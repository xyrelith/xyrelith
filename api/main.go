package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"xyrelith/api/endpoints"
	"xyrelith/api/models"
	"xyrelith/api/openai"
	"xyrelith/api/postgres"
)

func main() {
	x_postgres.Init()
	x_openai.Init()

	x_postgres.DB.AutoMigrate(models.Event{})

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/api/createEvent", endpoints.CreateEvent)
	router.GET("/api/scheduleEvents", endpoints.ScheduleEvents)
	router.GET("/api/listEvents", endpoints.ListEvents)
	router.Run(":2712")
}

package endpoints

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"

	"xyrelith/api/models"
	"xyrelith/api/openai"
	"xyrelith/api/postgres"
)

func ScheduleEvents(c *gin.Context) {
	var events []models.Event
	now := time.Now()
	ctx := context.Background()

	dbResult := x_postgres.DB.Where("start_date >= ?", now).Find(&events)
	if dbResult.Error != nil {
		log.Fatalf("query failed: %v", dbResult.Error)
	}

	eventsJSON, err := json.Marshal(events)
	if err != nil {
		log.Fatalf("Failed to marshal events: %v", err)
	}

	schema, err := jsonschema.GenerateSchemaForType(models.ScheduleResult{})
	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
	}

	resp, err := x_openai.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT5Nano,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `You are a AI calendar organizer that will help user automatically schedule tasks and events based on their time and name and priority
				Priority 1 is Critical and 4 Low. Make all events not collide u can move events to other days in the future.`,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: string(eventsJSON),
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   "ai_planner",
				Schema: schema,
				Strict: true,
			},
		},
	})

	if err != nil {
		log.Fatalf("CreateChatCompletion error: %v", err)
	}

	var result models.ScheduleResult
	err = schema.Unmarshal(resp.Choices[0].Message.Content, &result)
	if err != nil {
		log.Fatalf("Unmarshal schema error: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"events": result,
	})
}

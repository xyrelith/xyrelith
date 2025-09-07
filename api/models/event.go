package models

import "time"

type Priority int64

const (
	Low      Priority = 4
	Medium   Priority = 3
	High     Priority = 2
	Critical Priority = 1
)

type Event struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Priority  Priority  `json:"priority"`
}

type AIEvent struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Priority  Priority `json:"priority"`
}

type ScheduleResult struct {
	Reasoning string    `json:"reasoning"`
	Events    []AIEvent `json:"events"`
}

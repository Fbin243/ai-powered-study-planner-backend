package dto

import "time"

type TimeTrackDto struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"task_id" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	IsBreak   bool      `json:"is_break" validate:"omitempty,required"`
}

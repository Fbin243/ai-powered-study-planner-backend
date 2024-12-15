package dto

import "time"

type TaskDto struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name" validate:"required,min=1,max=100"`
	Description string    `json:"description,omitempty" validate:"omitempty,max=255"`
	Priority    string    `json:"priority,omitempty" validate:"required,oneof=low medium high"`
	Status      string    `json:"status,omitempty" validate:"required,oneof=not_started in_progress completed"`
	StartDate   time.Time `json:"start_date,omitempty" validate:"required"`
	EndDate     time.Time `json:"end_date,omitempty" validate:"required,gtfield=StartDate"`
}

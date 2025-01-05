package dto

import "time"

type DateTimeFilter struct {
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_date" validate:"omitempty,gtfield=StartDate"`
}

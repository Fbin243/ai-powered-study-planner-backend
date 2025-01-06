package dto

import "time"

type DateTimeFilter struct {
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time" validate:"omitempty,gtfield=StartTime"`
}

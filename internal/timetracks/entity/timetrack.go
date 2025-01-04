package entity

import (
	"ai-powered-study-planner-backend/pkg/db"
	"time"
)

type Timetrack struct {
	*db.BaseModel `bson:",inline"`
	FirebaseUID   string    `json:"firebase_uid" bson:"firebase_uid"`
	TaskID        string    `json:"task_id" bson:"task_id"`
	StartTime     time.Time `json:"start_time" bson:"start_time"`
	EndTime       time.Time `json:"end_time" bson:"end_time"`
	IsBreak       bool      `json:"is_break" bson:"is_break"`
}

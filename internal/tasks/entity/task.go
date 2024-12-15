package entity

import (
	"ai-powered-study-planner-backend/pkg/db"
	"time"
)

type Task struct {
	*db.BaseModel `bson:",inline"`
	FirebaseUID   string       `json:"firebase_uid,omitempty" bson:"firebase_uid"`
	Name          string       `json:"name" bson:"name"`
	Description   string       `json:"description,omitempty" bson:"description"`
	Priority      TaskPriority `json:"priority,omitempty" bson:"priority"`
	Status        TaskStatus   `json:"status,omitempty" bson:"status"`
	StartDate     time.Time    `json:"start_date,omitempty" bson:"start_date"`
	EndDate       time.Time    `json:"end_date,omitempty" bson:"end_date"`
}

type TaskPriority string

const (
	HighPriority   TaskPriority = "high"
	MediumPriority TaskPriority = "medium"
	LowPriority    TaskPriority = "low"
)

type TaskStatus string

const (
	NotStarted TaskStatus = "not_started"
	InProgress TaskStatus = "in_progress"
	Completed  TaskStatus = "completed"
	Expired    TaskStatus = "expired"
)

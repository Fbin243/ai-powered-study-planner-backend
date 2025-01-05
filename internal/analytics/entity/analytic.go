package entity

import "ai-powered-study-planner-backend/internal/timetracks/entity"

type Analytic struct {
	FirebaseUID        string             `json:"firebase_uid" bson:"firebase_uid"`
	TotalEstimatedTime int32              `json:"total_estimated_time" bson:"total_estimated_time"`
	TotalFocusTime     int32              `json:"total_focus_time" bson:"total_focus_time"`
	TotalBreakTime     int32              `json:"total_break_time" bson:"total_break_time"`
	TotalTasks         int32              `json:"total_tasks" bson:"total_tasks"`
	TaskDistribution   TaskDistribution   `json:"task_distribution" bson:"task_distribution"`
	Timetracks         []entity.Timetrack `json:"time_tracks" bson:"time_tracks"`
}

type TaskDistribution struct {
	NotStarted int32
	InProgress int32
	Completed  int32
	Expired    int32
}

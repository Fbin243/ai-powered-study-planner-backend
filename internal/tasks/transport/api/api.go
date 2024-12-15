package api

import "ai-powered-study-planner-backend/internal/tasks/business"

type TasksAPI struct {
	TasksBusiness *business.TasksBusiness
}

func NewTasksAPI(tasksBusiness *business.TasksBusiness) *TasksAPI {
	return &TasksAPI{
		TasksBusiness: tasksBusiness,
	}
}

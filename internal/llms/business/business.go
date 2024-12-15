package business

import (
	"ai-powered-study-planner-backend/internal/tasks/repo"
)

type LLMsBusiness struct {
	TasksRepo *repo.TasksRepo
}

func NewLLMsBusiness(tasksRepo *repo.TasksRepo) *LLMsBusiness {
	return &LLMsBusiness{
		TasksRepo: tasksRepo,
	}
}

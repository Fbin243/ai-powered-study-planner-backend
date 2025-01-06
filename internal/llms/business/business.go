package business

import (
	"ai-powered-study-planner-backend/internal/tasks/repo"
	timetracksRepo "ai-powered-study-planner-backend/internal/timetracks/repo"
)

type LLMsBusiness struct {
	TasksRepo      *repo.TasksRepo
	TimetracksRepo *timetracksRepo.TimetracksRepo
}

func NewLLMsBusiness(tasksRepo *repo.TasksRepo, timetracksRepo *timetracksRepo.TimetracksRepo) *LLMsBusiness {
	return &LLMsBusiness{
		TasksRepo:      tasksRepo,
		TimetracksRepo: timetracksRepo,
	}
}

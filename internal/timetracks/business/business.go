package business

import (
	"ai-powered-study-planner-backend/internal/timetracks/repo"
	tasksRepo "ai-powered-study-planner-backend/internal/tasks/repo"
)

type TimetracksBusiness struct {
	TimetracksRepo *repo.TimetracksRepo
	TasksRepo *tasksRepo.TasksRepo
}

func NewTimetracksBusiness(repo *repo.TimetracksRepo, tasksRepo *tasksRepo.TasksRepo) *TimetracksBusiness {
	return &TimetracksBusiness{repo, tasksRepo}
}

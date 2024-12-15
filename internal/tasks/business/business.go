package business

import "ai-powered-study-planner-backend/internal/tasks/repo"

type TasksBusiness struct {
	TasksRepo *repo.TasksRepo
}

func NewTasksBusiness(tasksRepo *repo.TasksRepo) *TasksBusiness {
	return &TasksBusiness{
		TasksRepo: tasksRepo,
	}
}

package composer

import (
	"ai-powered-study-planner-backend/internal/tasks/business"
	"ai-powered-study-planner-backend/internal/tasks/repo"
	"ai-powered-study-planner-backend/internal/tasks/transport/api"
	"ai-powered-study-planner-backend/pkg/db"
)

func ComposeTasksAPI() *api.TasksAPI {
	tasksRepo := repo.NewTasksRepo(db.New().GetCollection(db.TasksCollection))
	tasksBusiness := business.NewTasksBusiness(tasksRepo, db.GetRedisClient())

	return api.NewTasksAPI(tasksBusiness)
}

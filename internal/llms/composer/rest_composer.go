package composer

import (
	"ai-powered-study-planner-backend/internal/llms/business"
	"ai-powered-study-planner-backend/internal/llms/transport/api"
	"ai-powered-study-planner-backend/internal/tasks/repo"
	timetracksRepo "ai-powered-study-planner-backend/internal/timetracks/repo"
	"ai-powered-study-planner-backend/pkg/db"
)

func ComposeLLMsAPI() *api.LLMsAPI {
	tasksRepo := repo.NewTasksRepo(db.New().GetCollection(db.TasksCollection))
	timetracksRepo := timetracksRepo.NewTimetracksRepo(db.New().GetCollection(db.TimeTracksCollection))
	llmsBusiness := business.NewLLMsBusiness(tasksRepo, timetracksRepo, db.GetRedisClient())

	return api.NewLLMsAPI(llmsBusiness)
}

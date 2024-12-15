package composer

import (
	"ai-powered-study-planner-backend/internal/llms/business"
	"ai-powered-study-planner-backend/internal/llms/transport/api"
	"ai-powered-study-planner-backend/internal/tasks/repo"
	"ai-powered-study-planner-backend/pkg/db"
)

func ComposeLLMsAPI() *api.LLMsAPI {
	tasksRepo := repo.NewTasksRepo(db.New().GetCollection(db.TasksCollection))
	llmsBusiness := business.NewLLMsBusiness(tasksRepo)

	return api.NewLLMsAPI(llmsBusiness)
}

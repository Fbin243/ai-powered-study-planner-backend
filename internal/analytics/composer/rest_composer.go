package composer

import (
	"ai-powered-study-planner-backend/internal/analytics/business"
	"ai-powered-study-planner-backend/internal/analytics/transport/api"
	tasksRepo "ai-powered-study-planner-backend/internal/tasks/repo"
	timetracksRepo "ai-powered-study-planner-backend/internal/timetracks/repo"
	"ai-powered-study-planner-backend/pkg/db"
)

func ComposeAnalyticsAPI() *api.AnalyticsAPI {
	tasksRepo := tasksRepo.NewTasksRepo(db.New().GetCollection(db.TasksCollection))
	timetracksRepo := timetracksRepo.NewTimetracksRepo(db.New().GetCollection(db.TimeTracksCollection))
	redisClient := db.GetRedisClient()
	analyticsBusiness := business.NewAnalyticsBusiness(redisClient, timetracksRepo, tasksRepo)

	return api.NewAnalyticsAPI(analyticsBusiness)
}

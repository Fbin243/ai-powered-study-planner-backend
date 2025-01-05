package business

import (
	tasksRepo "ai-powered-study-planner-backend/internal/tasks/repo"
	"ai-powered-study-planner-backend/internal/timetracks/repo"

	"github.com/redis/go-redis/v9"
)

type AnalyticsBusiness struct {
	RedisClient    *redis.Client
	TimetracksRepo *repo.TimetracksRepo
	TasksRepo      *tasksRepo.TasksRepo
}

func NewAnalyticsBusiness(redisClient *redis.Client, repo *repo.TimetracksRepo, tasksRepo *tasksRepo.TasksRepo) *AnalyticsBusiness {
	return &AnalyticsBusiness{redisClient, repo, tasksRepo}
}

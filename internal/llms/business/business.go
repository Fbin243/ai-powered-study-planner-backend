package business

import (
	"ai-powered-study-planner-backend/internal/tasks/repo"
	timetracksRepo "ai-powered-study-planner-backend/internal/timetracks/repo"

	"github.com/redis/go-redis/v9"
)

type LLMsBusiness struct {
	TasksRepo      *repo.TasksRepo
	TimetracksRepo *timetracksRepo.TimetracksRepo
	RedisClient    *redis.Client
}

func NewLLMsBusiness(tasksRepo *repo.TasksRepo, timetracksRepo *timetracksRepo.TimetracksRepo, redisClient *redis.Client) *LLMsBusiness {
	return &LLMsBusiness{
		TasksRepo:      tasksRepo,
		TimetracksRepo: timetracksRepo,
		RedisClient:    redisClient,
	}
}

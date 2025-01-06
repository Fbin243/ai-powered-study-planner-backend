package business

import (
	tasksRepo "ai-powered-study-planner-backend/internal/tasks/repo"
	"ai-powered-study-planner-backend/internal/timetracks/repo"

	"github.com/redis/go-redis/v9"
)

type TimetracksBusiness struct {
	TimetracksRepo *repo.TimetracksRepo
	TasksRepo      *tasksRepo.TasksRepo
	RedisClient    *redis.Client
}

func NewTimetracksBusiness(repo *repo.TimetracksRepo, tasksRepo *tasksRepo.TasksRepo, redisClient *redis.Client) *TimetracksBusiness {
	return &TimetracksBusiness{repo, tasksRepo, redisClient}
}

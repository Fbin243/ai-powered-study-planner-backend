package business

import (
	"ai-powered-study-planner-backend/internal/tasks/repo"

	"github.com/redis/go-redis/v9"
)

type TasksBusiness struct {
	TasksRepo   *repo.TasksRepo
	RedisClient *redis.Client
}

func NewTasksBusiness(tasksRepo *repo.TasksRepo, redisClient *redis.Client) *TasksBusiness {
	return &TasksBusiness{
		TasksRepo:   tasksRepo,
		RedisClient: redisClient,
	}
}

package repo

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"ai-powered-study-planner-backend/pkg/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type TasksRepo struct {
	*db.BaseRepo[entity.Task]
}

func NewTasksRepo(collection *mongo.Collection) *TasksRepo {
	return &TasksRepo{
		db.NewBaseRepo[entity.Task](collection),
	}
}

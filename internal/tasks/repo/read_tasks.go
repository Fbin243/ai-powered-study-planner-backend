package repo

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *TasksRepo) GetTasksByFirebaseUID(firebaseUID string) ([]entity.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tasks := []entity.Task{}
	cursor, err := r.Collection.Find(ctx, bson.M{"firebase_uid": firebaseUID})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

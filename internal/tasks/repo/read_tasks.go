package repo

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type TaskFilter struct {
	StartTime *time.Time
	EndTime   *time.Time
	Status    *entity.TaskStatus
}

func (r *TasksRepo) GetTasksByFirebaseUID(firebaseUID string, taskFilter *TaskFilter) ([]entity.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"firebase_uid": firebaseUID}
	if taskFilter != nil {
		if taskFilter.StartTime != nil {
			filter["start_date"] = bson.M{"$gte": *taskFilter.StartTime}
		}
		if taskFilter.EndTime != nil {
			filter["end_date"] = bson.M{"$lte": *taskFilter.EndTime}
		}
		if taskFilter.Status != nil {
			filter["status"] = *taskFilter.Status
		}
	}

	tasks := []entity.Task{}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

package repo

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *TasksRepo) GetTasksByFirebaseUID(firebaseUID string, startTime *time.Time, endTime *time.Time) ([]entity.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"firebase_uid": firebaseUID}
	if startTime != nil {
		filter["start_time"] = bson.M{"$gte": *startTime}
	}
	if endTime != nil {
		filter["end_time"] = bson.M{"$lte": *endTime}
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

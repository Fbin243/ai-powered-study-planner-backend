package business

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/db"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *TasksBusiness) DeleteTask(ctx context.Context, taskID string) (*entity.Task, error) {
	// Check if the user is authenticated
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	// Validate taskId
	oid, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, errors.ErrBadRequest
	}

	// Check if the task exists
	task, err := b.TasksRepo.FindById(oid)
	if err != nil {
		return nil, errors.ErrBadRequest
	}

	// Check if the user is authorized to delete the task
	if task.FirebaseUID != firebaseProfile.UID {
		return nil, errors.ErrUserUnauthorized
	}

	// Invalidate caching
	err = b.RedisClient.Del(ctx, db.AIAnalyzeKey(firebaseProfile.UID)).Err()
	if err != nil {
		return nil, err
	}

	err = b.RedisClient.Del(ctx, db.AIFeedbackKey(firebaseProfile.UID)).Err()
	if err != nil {
		return nil, err
	}

	err = b.RedisClient.Del(ctx, db.AnalyticsKey(firebaseProfile.UID)).Err()
	if err != nil {
		return nil, err
	}

	return b.TasksRepo.DeleteById(oid)
}

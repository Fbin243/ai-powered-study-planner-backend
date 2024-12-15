package business

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *TasksBusiness) GetTasks(ctx context.Context) ([]entity.Task, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	return b.TasksRepo.GetTasksByFirebaseUID(firebaseProfile.UID)
}

func (b *TasksBusiness) GetTaskByID(ctx context.Context, taskID string) (*entity.Task, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	oid, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, errors.ErrBadRequest
	}

	task, err := b.TasksRepo.FindById(oid)
	if err != nil {
		return nil, err
	}

	if task.FirebaseUID != firebaseProfile.UID {
		return nil, errors.ErrUserUnauthorized
	}

	return task, nil
}

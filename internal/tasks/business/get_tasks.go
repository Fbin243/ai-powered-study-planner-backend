package business

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"
	"time"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *TasksBusiness) GetTasks(ctx context.Context) ([]entity.Task, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	tasks, err := b.TasksRepo.GetTasksByFirebaseUID(firebaseProfile.UID, nil, nil)
	if err != nil {
		return nil, err
	}

	// Keep the status of task up to current date
	tasks = lo.Map(tasks, func(task entity.Task, _ int) entity.Task {
		if task.Status != entity.Expired && task.Status != entity.Completed {
			currentTime := time.Now()
			needToUpdate := false
			if task.StartDate.Before(currentTime) &&
				task.EndDate.After(currentTime) &&
				task.Status != entity.InProgress {
				task.Status = entity.InProgress
				needToUpdate = true
			} else if task.EndDate.Before(currentTime) &&
				task.Status != entity.Expired {
				task.Status = entity.Expired
				needToUpdate = false
			}

			if needToUpdate {
				b.TasksRepo.UpdateById(task.ID, &task)
			}
		}

		return task
	})

	return tasks, nil
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

package business

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"ai-powered-study-planner-backend/internal/tasks/transport/dto"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/db"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *TasksBusiness) UpsertTask(ctx context.Context, task *dto.TaskDto) (*entity.Task, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	var upsertTask *entity.Task
	if task.ID != "" {
		oid, err := primitive.ObjectIDFromHex(task.ID)
		if err != nil {
			return nil, errors.ErrBadRequest
		}

		taskEntity, err := b.TasksRepo.FindById(oid)
		if err != nil {
			return nil, err
		}

		if taskEntity.FirebaseUID != firebaseProfile.UID {
			return nil, errors.ErrUserUnauthorized
		}

		taskEntity.Name = task.Name
		taskEntity.Description = task.Description
		taskEntity.Priority = entity.TaskPriority(task.Priority)
		taskEntity.Status = entity.TaskStatus(task.Status)
		taskEntity.EndDate = task.EndDate
		taskEntity.StartDate = task.StartDate

		upsertTask, err = b.TasksRepo.UpdateById(oid, taskEntity)
		if err != nil {
			return nil, err
		}
	} else {
		taskEntity := &entity.Task{
			BaseModel:   &db.BaseModel{},
			FirebaseUID: firebaseProfile.UID,
			Name:        task.Name,
			Description: task.Description,
			Priority:    entity.TaskPriority(task.Priority),
			Status:      entity.TaskStatus(task.Status),
			EndDate:     task.EndDate,
			StartDate:   task.StartDate,
		}

		var err error
		upsertTask, err = b.TasksRepo.Insert(taskEntity)
		if err != nil {
			return nil, err
		}
	}

	return upsertTask, nil
}

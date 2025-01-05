package business

import (
	taskEntity "ai-powered-study-planner-backend/internal/tasks/entity"
	"ai-powered-study-planner-backend/internal/timetracks/entity"
	"ai-powered-study-planner-backend/internal/timetracks/transport/dto"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/db"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (b *TimetracksBusiness) UpsertTimetrack(ctx context.Context, timetrack *dto.TimeTrackDto) (*entity.Timetrack, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	var upsertTimetrack *entity.Timetrack
	if timetrack.ID != "" {
		oid, err := primitive.ObjectIDFromHex(timetrack.ID)
		if err != nil {
			return nil, errors.ErrBadRequest
		}

		timetrackEntity, err := b.TimetracksRepo.FindById(oid)
		if err != nil {
			return nil, err
		}

		if timetrackEntity.FirebaseUID != firebaseProfile.UID {
			return nil, errors.ErrBadRequest
		}

		if !timetrackEntity.EndTime.IsZero() {
			return nil, fmt.Errorf("time track has already finished")
		}

		timetrackEntity.EndTime = time.Now()

		upsertTimetrack, err = b.TimetracksRepo.UpdateById(oid, timetrackEntity)
		if err != nil {
			return nil, err
		}
	} else {
		// Check difference between client and server
		serverTime := time.Now()
		if timetrack.StartTime.Add(20 * time.Second).Before(serverTime) {
			return nil, fmt.Errorf("insert session timeout")
		}

		// Check the status of task
		oid, err := primitive.ObjectIDFromHex(timetrack.TaskID)
		if err != nil {
			return nil, errors.ErrBadRequest
		}

		task, err := b.TasksRepo.FindById(oid)
		if err != nil {
			return nil, err
		}

		if task.FirebaseUID != firebaseProfile.UID {
			return nil, errors.ErrBadRequest
		}

		if task.Status != taskEntity.InProgress {
			return nil, fmt.Errorf("task is not in progress")
		}

		// Check if there is an active session
		currentTimetrack, err := b.TimetracksRepo.GetCurrentTimetrack(firebaseProfile.UID)
		if err != nil && err != mongo.ErrNoDocuments {
			return nil, err
		}

		if currentTimetrack != nil {
			return nil, fmt.Errorf("there is an current active session")
		}

		timetrackEntity := &entity.Timetrack{
			BaseModel:   &db.BaseModel{},
			FirebaseUID: firebaseProfile.UID,
			TaskID:      timetrack.TaskID,
			StartTime:   timetrack.StartTime,
			IsBreak:     timetrack.IsBreak,
		}

		upsertTimetrack, err = b.TimetracksRepo.Insert(timetrackEntity)
		if err != nil {
			return nil, err
		}
	}

	return upsertTimetrack, nil
}

func (b *TimetracksBusiness) GetCurrentTimetrack(ctx context.Context) (*entity.Timetrack, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	currentTimetrack, err := b.TimetracksRepo.GetCurrentTimetrack(firebaseProfile.UID)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return currentTimetrack, err
}

package business

import (
	"ai-powered-study-planner-backend/internal/analytics/transport/dto"
	"ai-powered-study-planner-backend/internal/tasks/entity"
	"ai-powered-study-planner-backend/internal/tasks/repo"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"
	"fmt"

	"github.com/samber/lo"
)

func (b *LLMsBusiness) GiveFeedbacks(ctx context.Context, dateTimeFilter *dto.DateTimeFilter) (*string, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	// Get all tasks of users from start time to end time
	tasks, err := b.TasksRepo.GetTasksByFirebaseUID(firebaseProfile.UID, &repo.TaskFilter{
		StartTime: dateTimeFilter.StartTime,
		EndTime:   dateTimeFilter.EndTime,
		Status:    lo.ToPtr(entity.NotStarted),
	})
	if err != nil {
		return nil, nil
	}

	// Get all timetracks of user from start time to end time
	fmt.Print(tasks)

	return nil, nil
}

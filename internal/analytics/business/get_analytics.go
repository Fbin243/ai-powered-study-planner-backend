package business

import (
	"ai-powered-study-planner-backend/internal/analytics/entity"
	"ai-powered-study-planner-backend/internal/analytics/transport/dto"
	tasksEntity "ai-powered-study-planner-backend/internal/tasks/entity"
	"ai-powered-study-planner-backend/internal/tasks/repo"
	timtracksEntity "ai-powered-study-planner-backend/internal/timetracks/entity"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/db"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

func (b *AnalyticsBusiness) GetAnalytics(ctx context.Context, dateFilter *dto.DateTimeFilter) (*entity.Analytic, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	analytics := &entity.Analytic{}

	// Get cache from redis
	analyticJSON, err := b.RedisClient.Get(ctx, db.AnalyticsKey(firebaseProfile.UID)).Result()
	if err == nil {
		err = json.Unmarshal([]byte(analyticJSON), analytics)
		if err != nil {
			return nil, err
		}

		return analytics, nil
	} else if err != redis.Nil {
		return nil, err
	}

	// Get all timetracks of user
	timetracks, err := b.TimetracksRepo.GetAllTimetracksOfProfile(firebaseProfile.UID, dateFilter.StartTime, dateFilter.EndTime)
	if err != nil {
		return nil, err
	}

	// Get all tasks of user
	tasks, err := b.TasksRepo.GetTasksByFirebaseUID(firebaseProfile.UID, &repo.TaskFilter{
		StartTime: dateFilter.StartTime,
		EndTime:   dateFilter.EndTime,
	})
	if err != nil {
		return nil, err
	}

	analytics.FirebaseUID = firebaseProfile.UID
	analytics.Timetracks = timetracks
	analytics.TotalTasks = int32(len(tasks))

	// Total focus time
	analytics.TotalFocusTime = lo.Reduce(timetracks, func(agg int32, timetrack timtracksEntity.Timetrack, _ int) int32 {
		if timetrack.IsBreak {
			return agg
		}

		return agg + int32(timetrack.EndTime.Sub(timetrack.StartTime).Seconds())
	}, 0)

	// Total break time
	analytics.TotalBreakTime = lo.Reduce(timetracks, func(agg int32, timetrack timtracksEntity.Timetrack, _ int) int32 {
		if !timetrack.IsBreak {
			return agg
		}

		return agg + int32(timetrack.EndTime.Sub(timetrack.StartTime).Seconds())
	}, 0)

	// Total estimated time
	analytics.TotalEstimatedTime = lo.Reduce(tasks, func(agg int32, task tasksEntity.Task, _ int) int32 {
		return agg + int32(task.EndDate.Sub(task.StartDate).Seconds())
	}, 0)

	// Task distribution
	for _, task := range tasks {
		switch task.Status {
		case tasksEntity.Completed:
			analytics.TaskDistribution.Completed++
		case tasksEntity.Expired:
			analytics.TaskDistribution.Expired++
		case tasksEntity.InProgress:
			analytics.TaskDistribution.InProgress++
		case tasksEntity.NotStarted:
			analytics.TaskDistribution.NotStarted++
		}
	}

	// Set analytics to redis
	JSON, err := json.Marshal(analytics)
	if err != nil {
		return nil, err
	}

	err = b.RedisClient.Set(ctx, db.AnalyticsKey(analytics.FirebaseUID), JSON, time.Hour).Err()
	if err != nil {
		return nil, err
	}

	fmt.Printf("redis")

	return analytics, nil
}

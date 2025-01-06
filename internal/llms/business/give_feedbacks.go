package business

import (
	"ai-powered-study-planner-backend/internal/analytics/transport/dto"
	"ai-powered-study-planner-backend/internal/tasks/repo"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/db"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func (b *LLMsBusiness) GiveFeedbacks(ctx context.Context, dateTimeFilter *dto.DateTimeFilter) (*string, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	// Get analyze result from redis
	cachedAnswer, err := b.RedisClient.Get(ctx, db.AIFeedbackKey(firebaseProfile.UID)).Result()
	if err == nil {
		return &cachedAnswer, nil
	} else if err != redis.Nil {
		return nil, err
	}

	// Get all tasks of users from start time to end time
	tasks, err := b.TasksRepo.GetTasksByFirebaseUID(firebaseProfile.UID, &repo.TaskFilter{
		StartTime: dateTimeFilter.StartTime,
		EndTime:   dateTimeFilter.EndTime,
	})
	if err != nil {
		return nil, nil
	}

	// Get all timetracks of user from start time to end time
	timetracks, err := b.TimetracksRepo.GetAllTimetracksOfProfile(firebaseProfile.UID, dateTimeFilter.StartTime, dateTimeFilter.EndTime)
	if err != nil {
		return nil, err
	}

	fmt.Printf("timetracks: %v\n", timetracks)
	fmt.Printf("tasks: %v\n", tasks)

	// Encode json and send to LLM to analyze
	data := map[string]interface{}{
		"current_date": time.Now(),
		"tasks":        tasks,
		"timetracks":   timetracks,
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling tasks: %v", err)
	}

	question := fmt.Sprintf(
		`Here is my task list and their timetracks (reference through task_id) in JSON format: %s. 
		Can you give me some feedbacks based on my focus and break time for my tasks?
		AI feedback:
		1. Identifying areas where I'm excelling.
		2. Suggesting subjects or tasks that may need more attention.
		3. Offering motivational feedback to encourage consistency and improvement.
		Please answer shortly under 300 tokens and return in Markdown format (Trim the markdown block code, just return the answer only.).`,
		string(dataJSON),
	)

	answer := b.Chat(question)
	if answer == nil {
		return nil, fmt.Errorf("error analyzing tasks, please try again later")
	}

	// Save answer to redis
	feedback, err := json.Marshal(answer)
	if err != nil {
		return nil, err
	}

	err = b.RedisClient.Set(ctx, db.AIFeedbackKey(firebaseProfile.UID), feedback, time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return answer, nil
}

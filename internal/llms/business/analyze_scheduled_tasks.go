package business

import (
	"ai-powered-study-planner-backend/internal/tasks/entity"
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
	"github.com/samber/lo"
)

func (b *LLMsBusiness) AnalyzeScheduledTasks(ctx context.Context) (*string, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	// Get analyze result from redis
	cachedAnswer, err := b.RedisClient.Get(ctx, db.AIAnalyzeKey(firebaseProfile.UID)).Result()
	if err == nil {
		return &cachedAnswer, nil
	} else if err != redis.Nil {
		return nil, err
	}

	// Get all tasks of users which have status not started
	tasks, err := b.TasksRepo.GetTasksByFirebaseUID(firebaseProfile.UID, &repo.TaskFilter{
		StartTime: lo.ToPtr(time.Now()),
		Status:    lo.ToPtr(entity.NotStarted),
	})
	if err != nil {
		return nil, nil
	}

	// Encode json and send to LLM to analyze
	data := map[string]interface{}{
		"current_date": time.Now(),
		"tasks":        tasks,
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling tasks: %v", err)
	}

	question := fmt.Sprintf(
		`Here is my task list in JSON format: %s. 
		Can you analyze my schedule and provide specific feedback for each task, focusing on:
		1. Any potential conflicts or overlaps in timing and how to resolve them.
		2. Recommendations for prioritizing tasks based on their urgency, importance, and current progress.
		3. Suggestions for improving time allocation and avoiding burnout for tasks with tight deadlines or long durations.
		4. Any other personalized tips to enhance my overall time management.
		Please answer concisely under 300 tokens and return in Markdown format (trim the block code, return the content only).`,
		string(dataJSON),
	)

	answer := b.Chat(question)
	if answer == nil {
		return nil, fmt.Errorf("error analyzing tasks, please try again later")
	}

	err = b.RedisClient.Set(ctx, db.AIAnalyzeKey(firebaseProfile.UID), answer, time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return answer, nil
}

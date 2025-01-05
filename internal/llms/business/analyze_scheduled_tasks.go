package business

import (
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/errors"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (b *LLMsBusiness) AnalyzeScheduledTasks(ctx context.Context) (*string, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	// Get all tasks of users
	tasks, err := b.TasksRepo.GetTasksByFirebaseUID(firebaseProfile.UID, nil, nil)
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
		Can you analyze my schedule and provide suggestions for improvements? Please include the following in your response:
		1. A warning if my schedule is too tight or may lead to burnout.
		2. Recommendations for prioritizing tasks to improve focus and balance.
		3. Any other suggestions to optimize my time management and avoid overloading.
		Please answer shortly under 300 tokens and return in Markdown format (Trim the markdown block code).`,
		string(dataJSON),
	)

	answer := b.Chat(question)
	if answer == nil {
		return nil, fmt.Errorf("error analyzing tasks, please try again later")
	}

	return answer, nil
}

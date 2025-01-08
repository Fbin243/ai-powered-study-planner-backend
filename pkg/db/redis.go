package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
	onceRedis   sync.Once
)

const (
	analytics   = "analytics_"
	ai_analyze  = "ai_analyze_"
	ai_feedback = "ai_feedback_"
)

func AnalyticsKey(firebaseUID string, startTime *time.Time, endTime *time.Time) string {
	key := analytics + firebaseUID
	if startTime != nil {
		key += startTime.String()
	}
	if endTime != nil {
		key += endTime.String()
	}

	return key
}

func AIAnalyzeKey(firebaseUID string) string {
	return ai_analyze + firebaseUID
}

func AIFeedbackKey(firebaseUID string, startTime *time.Time, endTime *time.Time) string {
	key := ai_feedback + firebaseUID
	if startTime != nil {
		key += startTime.String()
	}
	if endTime != nil {
		key += endTime.String()
	}

	return key
}

type AuthSession struct {
	ProfileID primitive.ObjectID
}

func GetRedisClient() *redis.Client {
	onceRedis.Do(func() {
		initRedis()
	})
	return redisClient
}

func initRedis() {
	redisURI := os.Getenv("REDIS_URI")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisURI == "" || redisPassword == "" {
		panic("REDIS_URI and REDIS_PASSWORD environment variables must be set")
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	log.Println("Redis connected:", pong)
}

func DeleteKeysByPattern(ctx context.Context, rdb *redis.Client, pattern string) error {
	var cursor uint64
	for {
		// Scan for keys matching the pattern
		keys, nextCursor, err := rdb.Scan(ctx, cursor, "*"+pattern+"*", 100).Result()
		if err != nil {
			return fmt.Errorf("error scanning keys: %w", err)
		}

		// Delete the matching keys
		if len(keys) > 0 {
			if err := rdb.Del(ctx, keys...).Err(); err != nil {
				return fmt.Errorf("error deleting keys: %w", err)
			}
			fmt.Printf("Deleted keys: %v\n", keys)
		}

		// Update the cursor for the next scan
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return nil
}

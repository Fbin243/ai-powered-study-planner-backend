package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

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

func AnalyticsKey(firebaseUID string) string {
	return analytics + firebaseUID
}

func AIAnalyzeKey(firebaseUID string) string {
	return ai_analyze + firebaseUID
}

func AIFeedbackKey(firebaseUID string) string {
	return ai_feedback + firebaseUID
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

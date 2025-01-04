package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	GetCollection(name string) *mongo.Collection
}

type service struct {
	db *mongo.Client
}

var (
	once     sync.Once
	instance *service
)

var (
	host     = os.Getenv("DB_HOST")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	database = os.Getenv("DB_DATABASE")
)

const (
	ProfilesCollection   = "profiles"
	TasksCollection      = "tasks"
	TimeTracksCollection = "time_tracks"
	AnalyticsCollection  = "analytics"
)

func New() Service {
	once.Do(func() {
		connectionString := fmt.Sprintf(
			"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
			username, password, host, database,
		)
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
		log.Printf("connection string: %s", connectionString)

		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		instance = &service{
			db: client,
		}
	})

	return instance
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetCollection(name string) *mongo.Collection {
	return s.db.Database(database).Collection(name)
}

package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"ai-powered-study-planner-backend/internal/profiles/business"
	"ai-powered-study-planner-backend/internal/profiles/repo"
	"ai-powered-study-planner-backend/pkg/db"
	"ai-powered-study-planner-backend/pkg/middlewares"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	db               db.Service
	authMiddleware   *middlewares.AuthMiddleware
	profilesBusiness *business.ProfileBusiness
}

func NewServer() *http.Server {
	// Dependency injection
	profilesRepo := repo.NewProfilesRepo(db.New().GetCollection(db.ProfilesCollection))
	profilesBusiness := business.NewProfileBusiness(profilesRepo)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db:               db.New(),
		authMiddleware:   middlewares.NewAuthMiddleware(),
		profilesBusiness: profilesBusiness,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

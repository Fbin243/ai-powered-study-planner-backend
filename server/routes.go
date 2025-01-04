package server

import (
	"net/http"

	llmsComposer "ai-powered-study-planner-backend/internal/llms/composer"
	"ai-powered-study-planner-backend/internal/profiles/composer"
	tasksComposer "ai-powered-study-planner-backend/internal/tasks/composer"
	timetracksComposer "ai-powered-study-planner-backend/internal/timetracks/composer"

	"ai-powered-study-planner-backend/pkg/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	authMiddleware := middlewares.NewAuthMiddleware()
	profilesApi := composer.ComposeProfilesAPI()
	tasksApi := tasksComposer.ComposeTasksAPI()
	llmsApi := llmsComposer.ComposeLLMsAPI()
	timetracksApi := timetracksComposer.ComposeTimetracksAPI()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.Use(authMiddleware.CheckAuth)
	// Profiles
	r.GET("/profile", profilesApi.GetProfile)
	r.POST("/profile", profilesApi.UpdateProfile)
	// Tasks
	r.GET("/tasks", tasksApi.GetTasks)
	r.GET("/tasks/:id", tasksApi.GetTask)
	r.POST("/tasks", tasksApi.UpsertTask)
	r.DELETE("/tasks/:id", tasksApi.DeleteTask)
	// LLMs
	r.GET("/llms/tasks", llmsApi.AnalyzeScheduledTasks)
	// Timetracks
	r.POST("/timetracks", timetracksApi.UpsertTimetrack)
	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

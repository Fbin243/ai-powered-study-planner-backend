package composer

import (
	tasksRepo "ai-powered-study-planner-backend/internal/tasks/repo"
	"ai-powered-study-planner-backend/internal/timetracks/business"
	"ai-powered-study-planner-backend/internal/timetracks/repo"
	"ai-powered-study-planner-backend/internal/timetracks/transport/api"
	"ai-powered-study-planner-backend/pkg/db"
)

func ComposeTimetracksAPI() *api.TimetracksAPI {
	timetracksRepo := repo.NewTimetracksRepo(db.New().GetCollection(db.TimeTracksCollection))
	tasksRepo := tasksRepo.NewTasksRepo(db.New().GetCollection(db.TasksCollection))
	timetracksBusiness := business.NewTimetracksBusiness(timetracksRepo, tasksRepo)

	return api.NewTimetracksAPI(timetracksBusiness)
}

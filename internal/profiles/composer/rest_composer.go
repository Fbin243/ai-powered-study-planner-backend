package composer

import (
	"ai-powered-study-planner-backend/internal/profiles/business"
	"ai-powered-study-planner-backend/internal/profiles/repo"
	"ai-powered-study-planner-backend/internal/profiles/transport/api"
	"ai-powered-study-planner-backend/pkg/db"
)

func ComposeProfilesAPI() *api.ProfilesAPI {
	profilesRepo := repo.NewProfilesRepo(db.New().GetCollection(db.ProfilesCollection))
	profilesBusiness := business.NewProfileBusiness(profilesRepo)

	return api.NewProfilesAPI(profilesBusiness)
}

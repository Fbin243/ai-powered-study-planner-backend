package business

import (
	"ai-powered-study-planner-backend/internal/profiles/repo"
)

type ProfilesBusiness struct {
	ProfilesRepo *repo.ProfilesRepo
}

func NewProfileBusiness(profilesRepo *repo.ProfilesRepo) *ProfilesBusiness {
	return &ProfilesBusiness{
		ProfilesRepo: profilesRepo,
	}
}

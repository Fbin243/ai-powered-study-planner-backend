package api

import "ai-powered-study-planner-backend/internal/profiles/business"

type ProfilesAPI struct {
	ProfileBusiness *business.ProfilesBusiness
}

func NewProfilesAPI(profileBusiness *business.ProfilesBusiness) *ProfilesAPI {
	return &ProfilesAPI{
		ProfileBusiness: profileBusiness,
	}
}

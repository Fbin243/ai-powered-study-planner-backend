package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *ProfilesAPI) GetProfile(c *gin.Context) {
	profile, err := api.ProfileBusiness.GetProfile(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

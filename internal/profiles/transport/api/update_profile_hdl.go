package api

import (
	"fmt"
	"net/http"

	"ai-powered-study-planner-backend/internal/profiles/transport/dto"
	"ai-powered-study-planner-backend/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (api *api) UpdateProfile(c *gin.Context) {
	var user dto.UpdateProfileDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Print("user: ", user)

	profile, err := api.ProfileBusiness.UpdateProfile(c.Request.Context(), &user)
	if err != nil {
		if err == errors.ErrUserUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

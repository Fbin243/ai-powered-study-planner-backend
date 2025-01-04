package api

import (
	"ai-powered-study-planner-backend/internal/timetracks/transport/dto"
	"ai-powered-study-planner-backend/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (api *TimetracksAPI) UpsertTimetrack(c *gin.Context) {
	var timetrack dto.TimeTrackDto
	if err := c.ShouldBindJSON(&timetrack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(timetrack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsertTimetrack, err := api.TimetracksBusiness.UpsertTimetrack(c.Request.Context(), &timetrack)
	if err != nil {
		if err == errors.ErrUserUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, upsertTimetrack)
}

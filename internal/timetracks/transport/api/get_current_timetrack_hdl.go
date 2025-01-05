package api

import (
	"ai-powered-study-planner-backend/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *TimetracksAPI) GetCurrentTimetrack(c *gin.Context) {
	timetrack, err := b.TimetracksBusiness.GetCurrentTimetrack(c.Request.Context())
	if err != nil {
		if err == errors.ErrUserUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, timetrack)
}

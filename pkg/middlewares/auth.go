package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"ai-powered-study-planner-backend/pkg/auth"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authManager *auth.FirebaseManager
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		authManager: auth.GetFirebaseManager(),
	}
}

func (m *AuthMiddleware) CheckAuth(c *gin.Context) {
	reqCtx := c.Request.Context()
	authKey := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(authKey, "Bearer ") {
		idToken := strings.Replace(authKey, "Bearer ", "", 1)
		firebaseProfile, err := auth.GetProfileByIDToken(idToken)
		if err != nil {
			log.Printf("invalid id token: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid id token"})
			return
		}

		reqCtx = context.WithValue(reqCtx, auth.ProfileKey, firebaseProfile)
	}

	c.Request = c.Request.WithContext(reqCtx)
	c.Next()
}

package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
)

type ClerkConfig struct {
	SecretKey string
}

type ClerkAdapter struct {
	Client clerk.Client
}

func NewClerkAdapter(config ClerkConfig) (AuthAdapter, error) {
	if config.SecretKey == "" {
		return nil, errors.New("clerk configuration missing: SecretKey is not set")
	}

	client, err := clerk.NewClient(config.SecretKey)
	if err != nil {
		return nil, err
	}

	return &ClerkAdapter{Client: client}, nil
}

func (adapter *ClerkAdapter) Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionToken := context.GetHeader("Authorization")
		if sessionToken == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
			return
		}

		if len(sessionToken) > 7 && sessionToken[:7] == "Bearer " {
			sessionToken = sessionToken[7:]
		}

		session, err := adapter.Client.VerifyToken(sessionToken)
		if err != nil {
			log.Printf("Clerk verification error: %v", err)
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		context.Set("auth_claims", session)
		context.Next()
	}
}
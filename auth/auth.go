package auth

import (
	"log"
	"net/http"
	"fmt"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
)

const ClaimsKey = "auth0_claims"

func AuthMiddleware(authConfig AuthConfig) gin.HandlerFunc {
	if authConfig.Domain == "" || authConfig.Audience == "" {
		panic("Auth configuration missing: Domain or Audience is not set")
	}

	jwkURL := fmt.Sprintf("https://%s/.well-known/jwks.json", authConfig.Domain)
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwkURL}, nil)
	configuration := auth0.NewConfiguration(client, []string{authConfig.Audience}, fmt.Sprintf("https://%s/", authConfig.Domain), "RS256")
	validator := auth0.NewValidator(configuration, nil)

	return func(context *gin.Context) {
		token, err := validator.ValidateRequest(context.Request)
		if err != nil {
			log.Printf("Auth error: %v", err)
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims := make(map[string]interface{})
		if err := token.Claims(&claims); err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "could not parse claims"})
			return
		}
		context.Set(ClaimsKey, claims)
		context.Next()
	}
}
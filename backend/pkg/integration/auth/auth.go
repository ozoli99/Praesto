package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
)

const ClaimsKey = "auth0_claims"

func AuthMiddleware() gin.HandlerFunc {
	domain := os.Getenv("AUTH0_DOMAIN")
	audience := os.Getenv("AUTH0_AUDIENCE")
	if domain == "" || audience == "" {
		panic("AUTH0_DOMAIN or AUTH0_AUDIENCE environment variable not set")
	}

	jwkURL := "https://" + domain + "/.well-known/jwks.json"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwkURL}, nil)
	configuration := auth0.NewConfiguration(client, []string{audience}, "https://" + domain + "/", "RS256")
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
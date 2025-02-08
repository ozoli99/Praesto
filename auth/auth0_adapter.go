package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
)

type Auth0Config struct {
	Domain   string
	Audience string
}

type Auth0Adapter struct {
	Configuration Auth0Config
}

func NewAuth0Adapter(config Auth0Config) AuthAdapter {
	return &Auth0Adapter{Configuration: config}
}

func (adapter *Auth0Adapter) Middleware() gin.HandlerFunc {
	if adapter.Configuration.Domain == "" || adapter.Configuration.Audience == "" {
		panic("Auth0 configuration missing: Domain or Audience is not set")
	}

	jwkURL := fmt.Sprintf("https://%s/.well-known/jwks.json", adapter.Configuration.Domain)
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwkURL}, nil)
	configuration := auth0.NewConfiguration(client, []string{adapter.Configuration.Audience}, fmt.Sprintf("https://%s/", adapter.Configuration.Domain), "RS256")
	validator := auth0.NewValidator(configuration, nil)

	return func(c *gin.Context) {
		token, err := validator.ValidateRequest(c.Request)
		if err != nil {
			log.Printf("Auth0 error: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims := make(map[string]interface{})
		if err := token.Claims(&claims); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "could not parse claims"})
			return
		}
		c.Set("auth_claims", claims)
		c.Next()
	}
}
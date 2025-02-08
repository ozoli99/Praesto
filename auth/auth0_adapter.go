package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Auth0Config struct {
	Domain       string
	ClientID     string
	ClientSecret string
	CallbackURL  string
	Audience     string
}

type Auth0Adapter struct {
	Provider     *oidc.Provider
	OAuth2Config oauth2.Config
	Verifier     *oidc.IDTokenVerifier
}

func NewAuth0Adapter(config Auth0Config) (AuthAdapter, error) {
	if config.Domain == "" || config.ClientID == "" || config.ClientSecret == "" || config.CallbackURL == "" || config.Audience == "" {
		return nil, errors.New("incomplete Auth0 configuration")
	}

	context := context.Background()
	provider, err := oidc.NewProvider(context, fmt.Sprintf("https://%s/", config.Domain))
	if err != nil {
		return nil, fmt.Errorf("failed to create OIDC provider: %w", err)
	}

	oauth2Config := oauth2.Config{
		ClientID: config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL: config.CallbackURL,
		Endpoint: provider.Endpoint(),
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: config.ClientID,
	})

	return &Auth0Adapter{
		Provider: provider,
		OAuth2Config: oauth2Config,
		Verifier: verifier,
	}, nil
}

func (adapter *Auth0Adapter) Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			return
		}

		tokenStr := parts[1]
		ctx := context.Request.Context()
		idToken, err := adapter.Verifier.Verify(ctx, tokenStr)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to parse token claims"})
			return
		}

		context.Set("auth_claims", claims)
		context.Next()
	}
}
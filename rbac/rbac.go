package rbac

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(context *gin.Context) {
		claimsInterface, exists := context.Get("auth0_claims")
		if !exists {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No claims found"})
			return
		}
		claims, ok := claimsInterface.(map[string]interface{})
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}
		namespace := "https://yourdomain.com"
		roles, ok := claims[namespace+"roles"].([]interface{})
		if !ok {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No roles in token"})
			return
		}
		authorized := false
		for _, r := range roles {
			if role, ok := r.(string); ok && role == requiredRole {
				authorized = true
				break
			}
		}
		if !authorized {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}
		context.Next()
	}
}
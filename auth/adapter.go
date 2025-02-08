package auth

import "github.com/gin-gonic/gin"

type AuthAdapter interface {
	Middleware() gin.HandlerFunc
}
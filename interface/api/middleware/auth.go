package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prrng/dealls/libs/auth"
)

// AuthMiddleware is a middleware for JWT authentication
type AuthMiddleware struct {
	jwtService *auth.JWTService
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware(jwtService *auth.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

// EmployeeAuth is a middleware to authenticate employees
func (m *AuthMiddleware) EmployeeAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		claims, err := m.jwtService.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		currentCtx := c.Request.Context()
		newCtx := context.WithValue(currentCtx, auth.CtxKeyAuthUserID, claims.ID)
		newCtx = context.WithValue(newCtx, auth.CtxKeyAuthUserEmail, claims.Email)
		newCtx = context.WithValue(newCtx, auth.CtxKeyAuthUserRole, claims.Role)

		c.Request = c.Request.WithContext(newCtx)

		c.Next()
	}
}

// AdminAuth is a middleware to authenticate admins
func (m *AuthMiddleware) AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		claims, err := m.jwtService.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		currentCtx := c.Request.Context()
		newCtx := context.WithValue(currentCtx, auth.CtxKeyAuthUserID, claims.ID)
		newCtx = context.WithValue(newCtx, auth.CtxKeyAuthUserEmail, claims.Email)
		newCtx = context.WithValue(newCtx, auth.CtxKeyAuthUserRole, claims.Role)

		c.Request = c.Request.WithContext(newCtx)

		c.Next()
	}
}

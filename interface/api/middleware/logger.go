package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prrng/dealls/domain/usecase"
)

// LoggerMiddleware is a middleware for logging HTTP requests
type LoggerMiddleware struct {
	auditUseCase usecase.AuditUseCase
}

// NewLoggerMiddleware creates a new instance of LoggerMiddleware
func NewLoggerMiddleware(auditUseCase usecase.AuditUseCase) *LoggerMiddleware {
	return &LoggerMiddleware{
		auditUseCase: auditUseCase,
	}
}

// Logger is a middleware to log HTTP requests
func (m *LoggerMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()
		
		// Log only if user is authenticated
		userID, exists := c.Get("user_id")
		if exists {
			// Get request details
			path := c.Request.URL.Path
			method := c.Request.Method
			statusCode := c.Writer.Status()
			clientIP := c.ClientIP()
			
			// Log request
			m.auditUseCase.LogAction(
				c.Request.Context(),
				userID.(uint),
				clientIP,
				method,
				path,
				0,
				"",
				"Status: "+string(rune(statusCode)),
			)
		}
	}
}

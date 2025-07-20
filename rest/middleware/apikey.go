package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pararang/hris/libs/httpresp"
)

type ApiKeyMiddleware struct {
	apiKey string
}

// NewApiKeyMiddleware creates a new instance of ApiKeyMiddleware
func NewApiKeyMiddleware(apiKey string) *ApiKeyMiddleware {
	return &ApiKeyMiddleware{
		apiKey: apiKey,
	}
}

func (m *ApiKeyMiddleware) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyHeader := c.GetHeader("x-api-key")
		if apiKeyHeader == "" || apiKeyHeader != m.apiKey {
			c.JSON(http.StatusUnauthorized, httpresp.Err(errors.New("unknown client")))
			c.Abort()
			return
		}

		c.Next()
	}
}

package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const CorrelationIDHeader = "X-Correlation-ID"

// CorrelationIDMiddleware generates or retrieves the correlation ID and attaches it to the context and response headers.
func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationID := c.GetHeader(CorrelationIDHeader)
		if correlationID == "" {
			correlationID = uuid.New().String()
		}

		c.Set(CorrelationIDHeader, correlationID)
		c.Writer.Header().Set(CorrelationIDHeader, correlationID)

		// set to context.Context
		ctx := context.WithValue(c.Request.Context(), CorrelationIDHeader, correlationID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

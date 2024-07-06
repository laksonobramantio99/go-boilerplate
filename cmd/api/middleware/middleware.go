package middleware

import (
	"context"
	"go-boilerplate/model/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CorrelationIDMiddleware generates or retrieves the correlation ID and attaches it to the context and response headers.
func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationID := c.GetHeader(constants.CORRELATION_ID)
		if correlationID == "" {
			correlationID = uuid.New().String()
		}

		c.Set(constants.CORRELATION_ID, correlationID)
		c.Writer.Header().Set(constants.CORRELATION_ID, correlationID)

		// set to context.Context
		ctx := context.WithValue(c.Request.Context(), constants.CORRELATION_ID, correlationID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

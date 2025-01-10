package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/logging"
)

// RequestIdMiddleware to generate and store requestId
func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate or fetch requestId
		requestID := c.GetHeader(ctx.RequestIdKey.String())
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Store requestId in Gin's context and set it globally
		c.Set(ctx.RequestIdKey.String(), requestID)
		requestCtx := context.WithValue(c.Request.Context(), ctx.RequestIdKey, requestID)
		c.Request = c.Request.WithContext(requestCtx)

		// Set the requestId in the response header
		c.Writer.Header().Set(ctx.RequestIdKey.String(), requestID)

		// Log the request start with requestId
		log := logging.LoggerWithRequestID(requestCtx)
		log.Debugf("Request ' %s ' started", c.Request.URL)

		// Process the request
		c.Next()

		// Clear the requestId after the request finishes
		log.Debugf("Request ' %s ' finished", c.Request.URL)
	}
}

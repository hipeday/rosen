package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/dto"
	"github.com/hipeday/rosen/internal/exception"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/hipeday/rosen/internal/messages"
	"net/http"
)

// ErrorHandlerMiddleware is a glob al middleware for handling exceptions
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			logger := logging.LoggerWithRequestID(c.Request.Context())
			// Recover from panic if one occurred
			if err := recover(); err != nil {
				// Log the error for debugging purposes
				logger.Errorf("Recovered from panic: %v\n", err)
				switch e := err.(type) {
				case exception.NotFoundError:
					c.JSON(e.Status(), dto.NewErrorResponse(messages.GetMessage(messages.DataDoesNotExist, c, e.Values...), c.Request.Context()))
				case exception.ValidationError:
					c.JSON(e.Status(), dto.NewErrorResponse(messages.GetMessage(e.Message, c), c.Request.Context()))
				case exception.UnauthorizedError:
					c.JSON(e.Status(), dto.NewErrorResponse(messages.GetMessage(messages.Unauthorized, c), c.Request.Context()))
				case exception.ForbiddenError:
					c.JSON(e.Status(), dto.NewErrorResponse(messages.GetMessage(e.Message, c), c.Request.Context()))
				case error:
					c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(e.Error(), c.Request.Context()))
				}
				// Prevent further handlers from running
				c.Abort()
			}
		}()
		// Proceed to the next middleware/handler
		c.Next()
	}
}

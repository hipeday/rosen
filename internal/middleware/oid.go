package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/dto"
	"github.com/hipeday/rosen/internal/messages"
	"net/http"
)

func OneIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		oid := c.GetHeader("One-Id")
		if oid == "" {
			c.JSON(http.StatusForbidden, dto.NewErrorResponse(messages.GetMessage(messages.OidCannotBeEmpty, c), c.Request.Context()))
			c.Abort()
		} else {
			// 将 OneId 存入上下文，供后续使用
			c.Set(ctx.OneIdKey.String(), oid)
			c.Next()
		}
	}
}

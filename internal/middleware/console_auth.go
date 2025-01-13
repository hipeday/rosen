package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/dto"
	"github.com/hipeday/rosen/internal/messages"
	"github.com/hipeday/rosen/internal/rdb"
	"github.com/hipeday/rosen/pkg/util/token"
	"net/http"
	"strings"
)

const tokenPrefix = "Bearer "

func ConsoleAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(messages.GetMessage(messages.Unauthorized, c), c.Request.Context()))
		c.Abort()
		return
	}

	tokenString, found := strings.CutPrefix(tokenString, tokenPrefix)
	if !found {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(messages.GetMessage(messages.Unauthorized, c), c.Request.Context()))
		c.Abort()
		return
	}

	// 验证 Token
	claims, err := token.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(messages.GetMessage(messages.Unauthorized, c), c.Request.Context()))
		c.Abort()
		return
	}

	// 从 Redis 检查 Token 是否有效
	cachedToken, err := ctx.GetRedisClient().Get(ctx.GetRedisContext(), rdb.ConsoleToken.String(claims.Username)).Result()
	if err != nil || cachedToken != tokenString {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(messages.GetMessage(messages.Unauthorized, c), c.Request.Context()))
		c.Abort()
		return
	}

	// 将 UserID 存入上下文，供后续使用
	c.Set(ctx.UsernameKey.String(), claims.Username)
	c.Next()
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/exception"
	"github.com/hipeday/rosen/pkg/util/token"
	"net/http"
)

func ConsoleAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		panic(exception.NewUnauthorizedError())
		return
	}

	// 验证 Token
	claims, err := token.ParseJWT(tokenString)
	if err != nil {
		panic(exception.NewUnauthorizedError())
		return
	}

	// 从 Redis 检查 Token 是否有效
	cachedToken, err := ctx.GetRedisClient().Get(ctx.GetRedisContext(), claims.Username).Result()
	if err != nil || cachedToken != tokenString {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired or invalid"})
		c.Abort()
		return
	}

	// 将 UserID 存入上下文，供后续使用
	c.Set("userID", claims.Username)
	c.Next()
}

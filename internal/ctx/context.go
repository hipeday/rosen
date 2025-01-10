package ctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/conf"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/redis/go-redis/v9"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

// GetLocalize 获取本次请求的国际化实例
func GetLocalize(c *gin.Context) *i18n.Localizer {
	return c.MustGet(LocalizeKey.String()).(*i18n.Localizer)
}

// GetRequestId 获取请求ID
func GetRequestId(ctx context.Context) (string, bool) {
	requestId, ok := ctx.Value(RequestIdKey).(string)
	return requestId, ok
}

// GetRedisContext 获取 redis 的 context
func GetRedisContext() context.Context {
	return ctx
}

func WithRedisClient(client *redis.Client) {
	redisClient = client
}

// GetRedisClient 获取 redis 客户端
func GetRedisClient() *redis.Client {
	return redisClient
}

// GetConfig 获取系统配置
func GetConfig() conf.Config {
	return conf.GetCfg()
}

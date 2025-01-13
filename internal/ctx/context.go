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

// GetOid 获取OneId
func GetOid(c *gin.Context) string {
	value, _ := c.Get(OneIdKey.String())
	return value.(string)
}

// GetUsername 获取当前登录用户的用户名
func GetUsername(c *gin.Context) string {
	return c.GetString(UsernameKey.String())
}

// GetRequestContext 获取请求上下文
func GetRequestContext(c *gin.Context) context.Context {
	return c.Request.Context()
}

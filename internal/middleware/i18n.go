package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/locales"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var localizeKey = "localize"

// I18nMiddleWare 一款 gin 的i18n本地化中间件 根据请求头分配当前语言类型
func I18nMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		langHeader := c.GetHeader("Accept-Language")
		if langHeader == "" {
			c.Header("Accept-Language", string(locales.DefaultLanguage))
		}
		// 设置 localize
		localize := i18n.NewLocalizer(locales.Bundle(), locales.MatchLanguage(langHeader))
		c.Set(localizeKey, localize)
		c.Next()
	}
}

func GetLocalize(c *gin.Context) *i18n.Localizer {
	return c.MustGet(localizeKey).(*i18n.Localizer)
}

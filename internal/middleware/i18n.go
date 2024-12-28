package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/locales"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strings"
)

var localizeKey = "localize"

// I18nMiddleWare 一款 gin 的i18n本地化中间件 根据请求头分配当前语言类型
func I18nMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		langHeader := c.GetHeader("Accept-Language")
		if langHeader == "" {
			c.Header("Accept-Language", string(locales.DefaultLanguage))
		}
		var lang = langHeader
		if strings.Contains(langHeader, ";") {
			langAndWeight := strings.Split(langHeader, ";")
			lang = langAndWeight[0]
			if strings.Contains(lang, ",") {
				lang = strings.Split(lang, ",")[0]
			}
		}
		// 设置 localize
		localize := i18n.NewLocalizer(locales.Bundle(), lang)
		c.Set(localizeKey, localize)
		c.Next()
	}
}

func GetLocalize(c *gin.Context) *i18n.Localizer {
	return c.MustGet(localizeKey).(*i18n.Localizer)
}

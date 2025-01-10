package ctx

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetLocalize(c *gin.Context) *i18n.Localizer {
	return c.MustGet(LocalizeKey.String()).(*i18n.Localizer)
}

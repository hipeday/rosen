package messages

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ErrorMessage string

func (e ErrorMessage) String() string {
	return string(e)
}

const defaultDictionaryPrefix = "dictionary"

func GetMessage(message ErrorMessage, c *gin.Context, values ...string) string {
	localize := ctx.GetLocalize(c)
	errorMessage := localize.MustLocalize(&i18n.LocalizeConfig{
		MessageID: string(message),
	})
	if values != nil {
		return fmt.Sprintf(errorMessage, getFormatMessage(c, values...)...)
	}
	return errorMessage
}

func getFormatMessage(c *gin.Context, values ...string) []any {
	localize := ctx.GetLocalize(c)
	if values == nil {
		return nil
	}
	var format []any
	for _, value := range values {
		format = append(format, localize.MustLocalize(&i18n.LocalizeConfig{
			MessageID: defaultDictionaryPrefix + "." + value,
		}))
	}
	return format
}

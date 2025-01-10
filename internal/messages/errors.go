package messages

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ErrorMessage string

const (
	WrongUsernameOrPassword      ErrorMessage = "messages.error.wrong_username_or_password" // 用户名或密码错误
	UsernameCannotBeEmpty        ErrorMessage = "messages.error.username_cannot_be_empty"   // 用户名或密码错误
	TOTPVerificationIsNotEnabled ErrorMessage = "messages.error.totp_verify_not_enabled"    // TOTP验证未启用
	TOTPCodeCannotBeEmpty        ErrorMessage = "messages.error.totp_code_cannot_be_empty"  // TOTP验证码不能为空

	DataDoesNotExist ErrorMessage = "messages.error.data_not_found" // 数据不存在
)

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

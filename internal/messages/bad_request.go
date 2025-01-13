package messages

const (
	WrongUsernameOrPassword      ErrorMessage = "messages.error.wrong_username_or_password"             // 用户名或密码错误
	UsernameCannotBeEmpty        ErrorMessage = "messages.error.username_cannot_be_empty"               // 用户名不能为空
	PasswordCannotBeEmpty        ErrorMessage = "messages.error.password_cannot_be_empty"               // 密码不能为空
	CaptchaCannotBeEmpty         ErrorMessage = "messages.error.captcha_cannot_be_empty"                // 验证码不能为空
	TOTPVerificationIsNotEnabled ErrorMessage = "messages.error.totp_verify_not_enabled"                // TOTP验证未启用
	TOTPCodeCannotBeEmpty        ErrorMessage = "messages.error.totp_code_cannot_be_empty"              // TOTP验证码不能为空
	TotpVerificationFailed       ErrorMessage = "messages.error.totp_verification_failed"               // totp验证失败
	CaptchaErrorOrDoesNotExist   ErrorMessage = "messages.error.the_captcha_is_wrong_or_does_not_exist" // 验证码错误或不存在 =
)

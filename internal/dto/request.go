package dto

// ConsoleLoginRequest 控制台管理员登录请求
type ConsoleLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
	OID      string
}

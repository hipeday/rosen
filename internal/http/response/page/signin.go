package page

import "github.com/hipeday/rosen/internal/http/response"

type SignInPageResponse struct {
	response.GeneralPageResponse        // 通用页面响应
	Description                  string `json:"description"`        // 系统描述
	SigninTitle                  string `json:"signin_title"`       // 登录标题
	Username                     string `json:"username"`           // 用户名
	Password                     string `json:"password"`           // 密码
	Captcha                      string `json:"captcha"`            // 验证码
	RememberMe                   string `json:"remember_me"`        // 记住我
	ForgotPassword               string `json:"forgot_password"`    // 忘记密码
	SignIn                       string `json:"signin"`             // 登录
	OtherSignIn                  string `json:"other_signin"`       // 其他登录方式
	SignupLinkBefore             string `json:"signup_link_before"` // 还没有账号？
	SignupLink                   string `json:"signin_link"`        // 注册
	About                        string `json:"about"`              // 关于
	PrivacyPolicy                string `json:"privacy_policy"`     // 隐私政策
	TermsOfService               string `json:"terms_of_service"`   // 服务条款
	ContactUs                    string `json:"contact_us"`         // 联系我们
	JoinUs                       string `json:"join_us"`            // 加入我们
}

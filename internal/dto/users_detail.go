package dto

type UsersDetail struct {
	Username      string `json:"username,omitempty"`        // 用户名
	Nickname      string `json:"nickname,omitempty"`        // 用户昵称
	RecentLoginAt int64  `json:"recent_login_at,omitempty"` // 最近登录时间
	Email         string `json:"email,omitempty"`           // 邮件地址
	SuperAdmin    bool   `json:"super_admin,omitempty"`     // 超级管理员
}

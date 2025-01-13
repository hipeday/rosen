package dto

import "github.com/hipeday/rosen/internal/repository"

type UsersDetail struct {
	Username      string        `json:"username,omitempty"`        // 用户名
	Nickname      string        `json:"nickname,omitempty"`        // 用户昵称
	RecentLoginAt int64         `json:"recent_login_at,omitempty"` // 最近登录时间
	Email         string        `json:"email,omitempty"`           // 邮件地址
	SuperAdmin    bool          `json:"super_admin,omitempty"`     // 超级管理员
	Profiles      UsersProfiles `json:"profiles,omitempty"`        // 用户配置信息
}

type UsersProfiles struct {
	Birthday     int64             `json:"birthday,omitempty"`      // 生日
	Gender       repository.Gender `json:"gender,omitempty"`        // 性别
	Address      string            `json:"address,omitempty"`       // 生日地址
	TOTPVerified bool              `json:"totp_verified,omitempty"` // 是否完成TOTP验证
	TOTPEnabled  bool              `json:"totp_enabled,omitempty"`  // 是否开启TOTP
	Mobile       string            `json:"mobile,omitempty"`        // 手机号
	Nickname     string            `json:"nickname,omitempty"`      // 用户昵称
	Avatar       string            `json:"avatar,omitempty"`        // 用户头像
	Bio          string            `json:"bio,omitempty"`           // 个人简介
}

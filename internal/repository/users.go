package repository

import (
	"github.com/hipeday/rosen/internal/logging"
	"time"
)

type UserStatus string

// Users 用户表实例
type Users struct {
	GeneralEntity[int64]
	Username    string     `db:"username"`      // 用户名 必须唯一
	Password    string     `db:"password"`      // 密码(md5哈希)
	TotpSecret  string     `db:"totp_secret"`   // TOTP 2FA的密钥
	Email       string     `db:"email"`         // 又想 必须唯一
	Status      UserStatus `db:"status"`        // 用户状态
	SuperAdmin  bool       `db:"super_admin"`   // 是否超级管理员，默认否
	LastLoginAt time.Time  `db:"last_login_at"` // 最后登录时间
}

// UsersRepository represents the users' repository.
type UsersRepository struct {
	DefaultRepository
}

// NewUsersRepository creates a new theme repository.
func NewUsersRepository() *UsersRepository {
	logger := logging.Logger()
	repository := UsersRepository{}
	defaultRepository, err := newRepositoryFactory(&repository)
	if err != nil {
		logger.Fatal(err)
	}
	return defaultRepository.(*UsersRepository)
}

func (r *UsersRepository) SelectByUsername(username string) (*Users, error) {
	connection := r.getConnection()
	var users Users
	err := connection.Get(&users, "SELECT * FROM users WHERE username = $1", username)
	if r.checkIsNoRowsErr(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UsersRepository) SelectByUsernameAndPassword(username, password string) (*Users, error) {
	connection := r.getConnection()
	var users Users
	err := connection.Get(&users, "SELECT * FROM users WHERE username = $1 AND password = $2", username, password)
	if r.checkIsNoRowsErr(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UsersRepository) UpdateById(users *Users) error {
	connection := r.getConnection()
	sql := "UPDATE rosen.public.users SET totp_secret = $1, email = $2, status = $3, super_admin = $4, last_login_at = $5, created_at = $6, modified_at = $7, created_by = $8, modified_by = $9 WHERE id = $10"
	_, err := connection.Exec(sql, users.TotpSecret, users.Email, users.Status, users.SuperAdmin, users.LastLoginAt, users.CreatedAt, time.Now(), 1, 1, users.Id)
	return err
}

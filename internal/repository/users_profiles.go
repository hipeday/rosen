package repository

import (
	"github.com/hipeday/rosen/internal/logging"
	"time"
)

type Gender string

const (
	Man    Gender = "M" // 男
	Female Gender = "F" // 女
	Other  Gender = "O" // 其他
)

// UsersProfiles 用户个人资料表
type UsersProfiles struct {
	GeneralEntity[int64]
	Userid       int64     `db:"userid"`        // 用户ID
	Birthday     time.Time `db:"birthday"`      // 生日
	Gender       Gender    `db:"gender"`        // 性别
	Address      string    `db:"address"`       // 地址
	TotpVerified bool      `db:"totp_verified"` // TOTP是否已经验证 默认否
	TotpEnabled  bool      `db:"totp_enabled"`  // TOTP是否开启
	Mobile       string    `db:"mobile"`        // 手机号码
	Nickname     string    `db:"nickname"`      // 昵称
	Avatar       string    `db:"avatar"`        // 头像
	Bio          string    `db:"bio"`           // 简介/自我描述
}

// UsersProfilesRepository represents the UsersProfilesRepository.
type UsersProfilesRepository struct {
	DefaultRepository
}

func (r *UsersProfilesRepository) SelectByUserid(userid int64) (*UsersProfiles, error) {
	var (
		connection    = r.getConnection()
		usersProfiles UsersProfiles
	)

	err := connection.Get(&usersProfiles, "SELECT * FROM rosen.public.users_profiles WHERE userid = $1", userid)
	if r.checkIsNoRowsErr(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &usersProfiles, nil
}

// NewUsersProfilesRepository creates a new UsersProfilesRepository.
func NewUsersProfilesRepository() *UsersProfilesRepository {
	logger := logging.Logger()
	repository := UsersProfilesRepository{}
	defaultRepository, err := newRepositoryFactory(&repository)
	if err != nil {
		logger.Fatal(err)
	}
	return defaultRepository.(*UsersProfilesRepository)
}

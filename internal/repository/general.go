package repository

import (
	"database/sql"
	"errors"
	"github.com/hipeday/rosen/internal/database"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/jmoiron/sqlx"
	"time"
)

// GeneralEntity 系统通用实例
type GeneralEntity[T any] struct {
	Id         T         `db:"id"`                             // 主键ID
	CreatedAt  time.Time `db:"created_at" json:"created_at"`   // 当前主题创建时间
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"` // 当前主题更新时间
	CreatedBy  int64     `db:"created_by" json:"created_by"`   // 当前主题创建人
	ModifiedBy int64     `db:"modified_by" json:"modified_by"` // 当前主题更新人
}

type Repository interface {
	checkIsNoRowsErr(err error) bool         // 检查异常是否是没有查到数据
	setConnection(connection *sqlx.DB) error // 设置连接
	getConnection() *sqlx.DB                 // 获取连接
}

type DefaultRepository struct {
	Repository
	db *sqlx.DB
}

func newRepositoryFactory(repo Repository) (Repository, error) {
	logger := logging.Logger()
	if repo == nil {
		return nil, errors.New("repository is nil")
	}
	db, err := database.Get()
	if err != nil {
		logger.Fatalf("get database connection error: %v", err)
	}
	err = repo.setConnection(db)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *DefaultRepository) checkIsNoRowsErr(err error) bool {
	return err != nil && errors.Is(err, sql.ErrNoRows)
}

func (r *DefaultRepository) setConnection(connection *sqlx.DB) error {
	if connection == nil {
		return errors.New("connection is nil")
	}
	r.db = connection
	return nil
}

func (r *DefaultRepository) getConnection() *sqlx.DB {
	return r.db
}

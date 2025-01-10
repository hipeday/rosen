package database

import (
	"fmt"
	"github.com/hipeday/rosen/conf"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/jmoiron/sqlx"
	"os"
)

var db *sqlx.DB

func initDatabase() error {
	cfg := ctx.GetConfig()
	switch cfg.Database.Type {
	case conf.PostgresSQL:
		return initPostgres()
	default:
		return fmt.Errorf("unsupported database driver: %s", cfg.Database.Type)
	}
}

// SetupDatabase 初始化数据库连接
func SetupDatabase() {
	if err := initDatabase(); db != nil && err != nil {
		logging.Logger().Fatal(err)
		os.Exit(1)
	}
	if err := Test(); err != nil {
		logging.Logger().Fatal(err)
		os.Exit(1)
	}
}

// Close 关闭数据库连接
func Close() {
	if db != nil {
		if err := db.Close(); err != nil {
			logging.Logger().Error(err)
		}
	}
}

// Get 获取数据库连接
func Get() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}
	return nil, fmt.Errorf("database connection is nil")
}

// Test 测试数据库连接
func Test() error {
	if db != nil {
		return db.Ping()
	}
	return fmt.Errorf("database connection is nil")
}

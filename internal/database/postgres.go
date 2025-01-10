package database

import (
	"fmt"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/logging"
	_ "github.com/jackc/pgx/v5/stdlib" // 注册 pgx 驱动
	"github.com/jmoiron/sqlx"
	"time"
)

// initPostgres initializes the PostgresSQL connection pool
func initPostgres() error {
	var (
		cfg              = ctx.GetConfig()
		database         = cfg.Database
		databasePoolConf = database.Pool
		err              error
	)

	// 数据库连接信息 username password host port dbname
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable options='-c search_path=%s'",
		database.Host,
		database.Port,
		database.Username,
		database.Password,
		database.Database,
		"public",
	)

	db, err = sqlx.Open("pgx", dsn)

	if err != nil {
		return err
	}

	// Set up database connection pool
	db.SetMaxOpenConns(databasePoolConf.MaxConn)
	db.SetMaxIdleConns(databasePoolConf.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(databasePoolConf.MaxLife) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(databasePoolConf.MaxIdleTime) * time.Hour)

	// Ping the database to verify the connection
	if err = db.Ping(); err != nil {
		return err
	}
	logging.Logger().Info("Database connection pool initialized.")
	return nil
}

package config

import (
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/hipeday/rosen/pkg/env"
	"os"
	"time"
)

func LoadConfiguration() {

	// 加载配置文件
	cfg := ctx.GetConfig()

	if env.Version() == "" {
		err := os.Setenv("ROSEN_VERSION", cfg.Application.Version)
		if err != nil {
			panic(err)
		}
	}

	if cfg.Application.TimeZone != "" {
		// 指定时区
		location, err := time.LoadLocation(cfg.Application.TimeZone)
		if err != nil {
			logging.Logger().Fatalf(err.Error(), err)
		}
		time.Local = location
	}

}

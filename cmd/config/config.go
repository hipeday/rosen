package config

import (
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/pkg/env"
	"os"
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
}

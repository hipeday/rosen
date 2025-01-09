package log

import (
	"github.com/hipeday/rosen/conf"
	"github.com/hipeday/rosen/internal/logging"
)

func Setup() {
	cfg := conf.GetCfg()
	logging.Setup(cfg.Logger)
}

package log

import (
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/logging"
)

func Setup() {
	cfg := ctx.GetConfig()
	logging.Setup(cfg.Logger)
}

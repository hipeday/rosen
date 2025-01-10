package page

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/ctx"
	"path/filepath"
)

// Login 登录页面
func Login(engine *gin.Engine) {

	cfg := ctx.GetConfig()

	consoleGroup := engine.Group("")
	consoleGroup.GET("/login", func(c *gin.Context) {
		file := filepath.Join(cfg.Theme.Path, cfg.Theme.Default, c.DefaultQuery("path", "index.html"))
		c.File(file)
	})
}

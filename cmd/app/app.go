package app

import (
	"github.com/hipeday/rosen/cmd/banner"
	"github.com/hipeday/rosen/cmd/config"
	"github.com/hipeday/rosen/cmd/hooks"
	"github.com/hipeday/rosen/cmd/log"
	"github.com/hipeday/rosen/cmd/route"
	"github.com/hipeday/rosen/internal/database"
)

//r.Static("/_next/static", rootPath+"/_next/static")
//r.LoadHTMLGlob(rootPath + "/*.html")

func Run() {
	config.LoadConfiguration()
	log.Setup()
	banner.MakeBanner()
	database.Setup()
	hooks.BindingHooks()

	route.Startup() // 启动路由服务器
}

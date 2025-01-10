package app

import (
	"github.com/hipeday/rosen/cmd/banner"
	"github.com/hipeday/rosen/cmd/config"
	"github.com/hipeday/rosen/cmd/hooks"
	"github.com/hipeday/rosen/cmd/log"
	"github.com/hipeday/rosen/cmd/route"
	"github.com/hipeday/rosen/internal/database"
	"github.com/hipeday/rosen/internal/rdb"
)

//r.Static("/_next/static", rootPath+"/_next/static")
//r.LoadHTMLGlob(rootPath + "/*.html")

func Run() {
	config.LoadConfiguration() // 初始化配置文件
	log.Setup()                // 初始化日志
	banner.MakeBanner()        // 初始化 banner
	database.SetupDatabase()   // 初始化 db
	rdb.SetupRedis()           // 初始化Redis
	hooks.BindingHooks()       // 绑定钩子
	route.Startup()            // 启动路由服务器
}

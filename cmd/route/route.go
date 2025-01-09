package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/cmd/route/console"
	"github.com/hipeday/rosen/cmd/route/sso/page"
	"github.com/hipeday/rosen/conf"
	"github.com/hipeday/rosen/internal/handler"
	"path/filepath"
	"strconv"
)

var (
	engine         *gin.Engine
	handlerFactory *handler.Factory
)

func Startup() {
	cfg := conf.GetCfg()
	server := cfg.Server
	gin.SetMode(server.Mode)
	engine = gin.Default()

	handlerFactory = handler.NewHandlerFactory()

	// init pages
	initPages()

	// init api
	initApi()

	addr := server.IP + ":" + strconv.Itoa(int(server.Port))
	err := engine.Run(addr)
	if err != nil {
		panic(err)
	}
}

// GetEngine 获取gin引擎
func GetEngine() (*gin.Engine, error) {
	if engine != nil {
		return engine, nil
	}
	return nil, fmt.Errorf("gin engine is not started")
}

func initPages() {
	cfg := conf.GetCfg()

	engine.Static("/_next/static", filepath.Join(cfg.Theme.Path, cfg.Theme.Default, "_next/static"))

	page.Login(engine)
}

func initApi() {
	// init console api
	console.InitConsoleApi(handlerFactory, engine.Group("/api/console"))
}

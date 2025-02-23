package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/cmd/route/console"
	"github.com/hipeday/rosen/cmd/route/sso/page"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/handler"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/hipeday/rosen/internal/middleware"
	"path/filepath"
	"strconv"
)

var (
	engine         *gin.Engine
	handlerFactory *handler.Factory
)

func Startup() {
	cfg := ctx.GetConfig()
	server := cfg.Server
	gin.SetMode(server.Mode)
	engine = gin.Default()
	engine.Use(middleware.ErrorHandlerMiddleware(), middleware.I18nMiddleWare(), middleware.RequestIdMiddleware(), middleware.OneIdMiddleware())

	handlerFactory = handler.NewHandlerFactory()

	// init pages
	initPages()

	// init api
	initApi()

	addr := server.IP + ":" + strconv.Itoa(int(server.Port))
	logging.Logger().Infof("Server is starting... Listening on http://%s:%d", server.IP, server.Port)
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
	cfg := ctx.GetConfig()

	engine.Static("/_next/static", filepath.Join(cfg.Theme.Path, cfg.Theme.Default, "_next/static"))

	page.Login(engine)
}

func initApi() {
	// init console api
	console.InitConsoleApi(handlerFactory, engine.Group("/api/console"))
}

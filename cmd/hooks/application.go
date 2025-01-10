package hooks

import (
	"github.com/hipeday/rosen/internal/database"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/hipeday/rosen/internal/rdb"
	"os"
	"os/signal"
	"syscall"
)

// BindingHooks 绑定系统钩子 该函数会进行一些例如:
// 系统退出
// 还需要其他系统钩子则在考虑添加。钩子逻辑采用 chan 交互
func BindingHooks() {
	destroyHook() // 系统退出hook钩子
}

func destroyHook() {
	signals := make(chan os.Signal, 1)

	// 监听中断信号和终止信号
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 启动 goroutine 监听信号
	go func() {
		logger := logging.Logger()
		sig := <-signals
		logger.Infof("Caught signal %s. Shutting down...", sig)
		database.Close()
		rdb.Close()
		logger.Infof("Safe shutdown successful")
		os.Exit(0)
	}()
}

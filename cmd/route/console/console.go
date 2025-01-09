package console

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/handler"
	"github.com/hipeday/rosen/internal/logging"
)

func InitConsoleApi(handlerFactory *handler.Factory, routeGroup *gin.RouterGroup) {
	log := logging.Logger()
	instance, err := handlerFactory.GetHandler(handler.Console)
	if err != nil {
		log.Fatal(err)
	}
	consoleHandler := instance.(*handler.ConsoleHandler)
	routeGroup.POST("/login", consoleHandler.Login)      // 控制台管理员登录
	routeGroup.POST("/logout", consoleHandler.Logout)    // 控制台管理员登出
	routeGroup.GET("/current", consoleHandler.Current)   // 获取当前登录的管理员信息
	routeGroup.GET("/setup2fa", consoleHandler.Setup2fa) // 获取2fa信息
}

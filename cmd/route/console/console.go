package console

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/handler"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/hipeday/rosen/internal/middleware"
)

func InitConsoleApi(handlerFactory *handler.Factory, routeGroup *gin.RouterGroup) {
	log := logging.Logger()
	instance, err := handlerFactory.GetHandler(handler.Console)
	if err != nil {
		log.Fatal(err)
	}
	consoleHandler := instance.(*handler.ConsoleHandler)

	whitelistRoute := routeGroup.Group("")
	whitelistRoute.POST("/login", consoleHandler.Login)    // 控制台管理员登录
	whitelistRoute.GET("/captcha", consoleHandler.Captcha) // 获取控制台管理员登录验证码

	consoleAuthRoute := routeGroup.Group("", middleware.ConsoleAuthMiddleware)
	consoleAuthRoute.POST("/logout", consoleHandler.Logout)  // 控制台管理员登出
	consoleAuthRoute.GET("/current", consoleHandler.Current) // 获取当前登录的管理员信息
	consoleAuthRoute.GET("/totp", consoleHandler.GetTOTP)    // 获取TOTP信息
	consoleAuthRoute.PUT("/totp", consoleHandler.Setup2fa)   // 设置TOTP二步验证
}

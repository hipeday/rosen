package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/dto"
	"github.com/hipeday/rosen/internal/exception"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/hipeday/rosen/internal/messages"
	"github.com/hipeday/rosen/internal/repository"
	"github.com/pquerna/otp/totp"
	"net/http"
)

type ConsoleHandler struct {
	usersRepo        *repository.UsersRepository
	userProfilesRepo *repository.UsersProfilesRepository
}

// Login 后台管理系统登录
func (c *ConsoleHandler) Login(ctx *gin.Context) {
	var (
		request dto.ConsoleLoginRequest
		err     error
	)

	if err = ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Rosen Console Admin Panel",
		AccountName: request.Username,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if request.Username == "rosen@hipeday.org" {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: messages.GetMessage(messages.WrongUsernameOrPassword, ctx)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"secret": key.Secret(),
		"qrcode": key.URL(),
	})

}

// Logout 后台管理系统登录
func (c *ConsoleHandler) Logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ping": "pong", "action": "logout"})
}

// Current 后台管理系统登录
func (c *ConsoleHandler) Current(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ping": "pong", "action": "logout"})
}

// Setup2fa 后台管理系统登录
func (c *ConsoleHandler) Setup2fa(ctx *gin.Context) {

	var (
		usersRepo = c.usersRepo
		err       error
		log       = logging.LoggerWithRequestID(ctx.Request.Context())
	)

	username := ctx.Query("username")  // 用户名
	totpCode := ctx.Param("totp_code") // 前一份2fa的验证码

	log.Debugf("API '%s' username: '%s', totpCode '%s'", ctx.Request.URL.Path, username, totpCode)

	if username == "" {
		panic(exception.NewValidationError(messages.UsernameCannotBeEmpty))
		return
	}

	users, err := usersRepo.SelectByUsername(username)
	if err != nil {
		panic(err)
		return
	}

	if users == nil {
		log.Debugf("Username %s does not exist from database", username)
		panic(exception.NewNotFoundError("users"))
		return
	}

	usersProfiles, err := c.userProfilesRepo.SelectByUserid(users.Id)
	if err != nil {
		log.Debugf("Query %s's profiles error", username)
		logging.Logger().Errorf(err.Error(), err)
		panic(err)
		return
	}

	if usersProfiles == nil {
		log.Debugf("%s's profiles not found from database", username)
		panic(exception.NewNotFoundError("users_profiles"))
		return
	}

	if !usersProfiles.TotpEnabled {
		log.Debugf("%s's totp is not enabled", username)
		panic(exception.NewValidationError(messages.TOTPVerificationIsNotEnabled))
		return
	}

	if usersProfiles.TotpVerified {
		if totpCode == "" {
			log.Debugf("%s's totp is verified. Please submit the totp cacptcha and try again。", username)
			panic(exception.NewValidationError(messages.TOTPCodeCannotBeEmpty))
			return
		}
		// TODO 校验TOTP code
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "RosenConsoleAdminPanel",
		AccountName: username,
	})

	if err != nil {
		panic(err)
		return
	}

	secret := key.Secret()

	users.TotpSecret = secret
	// 更新用户信息
	err = usersRepo.UpdateById(users)
	if err != nil {
		panic(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"secret": secret,
		"qrcode": key.URL(),
	})

}

// VerifyTOTP 验证 TOTP
//func VerifyTOTP(c *gin.Context) {
//	// 获取用户名和验证码
//	username := c.Query("username")
//	totpCode := c.Query("totp_code")
//
//	// 检查用户是否存在
//	secret, exists := userSecrets[username]
//	if !exists {
//		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户不存在"})
//		return
//	}
//
//	// 验证 TOTP 验证码
//	valid := totp.Validate(totpCode, secret)
//	if valid {
//		c.JSON(http.StatusOK, gin.H{"message": "2FA 验证成功"})
//	} else {
//		c.JSON(http.StatusUnauthorized, gin.H{"message": "2FA 验证失败"})
//	}
//}

func (c *ConsoleHandler) GetType() Type {
	return Console
}

func NewConsoleHandler() *ConsoleHandler {
	return &ConsoleHandler{
		usersRepo:        repository.NewUsersRepository(),
		userProfilesRepo: repository.NewUsersProfilesRepository(),
	}
}

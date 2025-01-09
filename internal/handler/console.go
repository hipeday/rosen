package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/dto"
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
	)

	username := ctx.Query("username")  // 用户名
	totpCode := ctx.Param("totp_code") // 前一份2fa的验证码

	if username == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: messages.GetMessage(messages.UsernameCannotBeEmpty, ctx),
		})
		return
	}

	users, err := usersRepo.SelectByUsername(username)
	if err != nil {
		logging.Logger().Errorf(err.Error(), err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if users == nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Error: messages.GetMessage(messages.DataDoesNotExist, ctx, "users")})
		return
	}

	usersProfiles, err := c.userProfilesRepo.SelectByUserid(users.Id)
	if err != nil {
		logging.Logger().Errorf(err.Error(), err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if usersProfiles == nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Error: messages.GetMessage(messages.DataDoesNotExist, ctx, "users_profiles")})
		return
	}

	if !usersProfiles.TotpEnabled {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: messages.GetMessage(messages.TOTPVerificationIsNotEnabled, ctx)})
		return
	}

	if usersProfiles.TotpVerified {
		if totpCode == "" {
			ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: messages.GetMessage(messages.TOTPCodeCannotBeEmpty, ctx)})
			return
		}
		// TODO 校验TOTP code
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "RosenConsoleAdminPanel",
		AccountName: username,
	})

	if err != nil {
		logging.Logger().Errorf(err.Error(), err)
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	secret := key.Secret()

	users.TotpSecret = secret
	// 更新用户信息
	err = usersRepo.UpdateById(users)
	if err != nil {
		logging.Logger().Errorf(err.Error(), err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
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

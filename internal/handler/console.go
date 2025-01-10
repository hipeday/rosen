package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/dto"
	"github.com/hipeday/rosen/internal/exception"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/hipeday/rosen/internal/messages"
	"github.com/hipeday/rosen/internal/rdb"
	"github.com/hipeday/rosen/internal/repository"
	"github.com/mojocn/base64Captcha"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"net/http"
	"time"
)

var store = base64Captcha.DefaultMemStore

type ConsoleHandler struct {
	usersRepo        *repository.UsersRepository
	userProfilesRepo *repository.UsersProfilesRepository
}

// Login 后台管理系统登录
func (h *ConsoleHandler) Login(ctx *gin.Context) {
	var (
		request dto.ConsoleLoginRequest
		err     error
	)

	if err = ctx.ShouldBindBodyWithJSON(&request); err != nil {
		panic(err)
		return
	}

}

func (h *ConsoleHandler) Captcha(c *gin.Context) {
	var (
		err error
		log = logging.LoggerWithRequestID(c.Request.Context())
	)
	oneId := c.Query("oid")
	if oneId == "" {
		panic(exception.NewValidationError(messages.OidCannotBeEmpty))
		return
	}

	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80) // 长、宽、位数、噪点密度、字体大小
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		panic(err)
		return
	}
	log.Debugf("API '%s' generate captcha is '%s'", c.Request.URL, answer)
	client := ctx.GetRedisClient()
	err = client.Set(ctx.GetRedisContext(), rdb.ConsoleLoginCaptcha.String(oneId), answer, 5*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "captcha": b64s})
}

// Logout 后台管理系统登录
func (h *ConsoleHandler) Logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ping": "pong", "action": "logout"})
}

// Current 后台管理系统登录
func (h *ConsoleHandler) Current(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ping": "pong", "action": "logout"})
}

func (h *ConsoleHandler) GetTOTP(ctx *gin.Context) {
	var (
		usersRepo = h.usersRepo
		key       *otp.Key
		err       error
		log       = logging.LoggerWithRequestID(ctx.Request.Context())
	)

	username := ctx.Query("username") // 用户名
	log.Debugf("API '%s' username: '%s'", ctx.Request.URL.Path, username)

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

	if users.TotpSecret != nil {
		key, err = totp.Generate(totp.GenerateOpts{
			Issuer:      "RosenConsoleAdminPanel",
			AccountName: username,
			Secret:      []byte(*users.TotpSecret),
		})
	} else {
		key, err = totp.Generate(totp.GenerateOpts{
			Issuer:      "RosenConsoleAdminPanel",
			AccountName: username,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"secret": key.Secret(),
		"qrcode": key.URL(),
	})

}

// Setup2fa 后台管理系统登录
func (h *ConsoleHandler) Setup2fa(ctx *gin.Context) {

	var (
		usersRepo = h.usersRepo
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

	usersProfiles, err := h.userProfilesRepo.SelectByUserid(users.Id)
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
		if validate := totp.Validate(totpCode, *users.TotpSecret); !validate {
			panic(exception.NewValidationError(messages.TotpVerificationFailed))
			return
		}
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

	users.TotpSecret = &secret

	usersProfiles.TotpVerified = false
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

func (h *ConsoleHandler) GetType() Type {
	return Console
}

func NewConsoleHandler() *ConsoleHandler {
	return &ConsoleHandler{
		usersRepo:        repository.NewUsersRepository(),
		userProfilesRepo: repository.NewUsersProfilesRepository(),
	}
}

package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/internal/dto"
	"github.com/hipeday/rosen/internal/exception"
	"github.com/hipeday/rosen/internal/logging"
	"github.com/hipeday/rosen/internal/messages"
	"github.com/hipeday/rosen/internal/rdb"
	"github.com/hipeday/rosen/internal/repository"
	"github.com/hipeday/rosen/pkg/util/token"
	"github.com/mojocn/base64Captcha"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var store = base64Captcha.DefaultMemStore

type ConsoleHandler struct {
	usersRepo        *repository.UsersRepository
	userProfilesRepo *repository.UsersProfilesRepository
}

// Login 后台管理系统登录
func (h *ConsoleHandler) Login(c *gin.Context) {
	var (
		request   dto.ConsoleLoginRequest
		err       error
		log       = logging.LoggerWithRequestID(c.Request.Context())
		usersRepo = h.usersRepo
	)

	if err = c.ShouldBindBodyWithJSON(&request); err != nil {
		panic(err)
		return
	}

	// 设置One Id
	request.OID = ctx.GetOid(c)

	h.loginValidation(request, log)

	log.Debugf("Request username: ' %s ', password: ' %s ', captcha: ' %s '", request.Username, request.Password, request.Captcha)

	users, err := usersRepo.SelectByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		panic(err)
		return
	}

	if users == nil {
		panic(exception.NewValidationError(messages.WrongUsernameOrPassword))
	}

	// 登录成功 1. 生成access_token 2. 更新最后的登录时间
	// token 5天
	expiresAt := time.Now().Add(7200 * time.Minute)
	accessToken, err := token.GenerateAdminPanelJWT(users.Username, ctx.GetConfig().Application.Name, expiresAt)
	if err != nil {
		panic(err)
	}

	users.LastLoginAt = time.Now()
	err = usersRepo.UpdateById(users)
	if err != nil {
		panic(err)
	}

	redisClient := ctx.GetRedisClient()
	redisClient.Set(ctx.GetRedisContext(), rdb.ConsoleToken.String(users.Username), accessToken, 7200*time.Minute)

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "expires_at": 7200})

}

func (h *ConsoleHandler) loginValidation(request dto.ConsoleLoginRequest, log *zap.SugaredLogger) {
	if request.Username == "" {
		log.Debugf("Username is empty")
		panic(exception.NewValidationError(messages.UsernameCannotBeEmpty))
		return
	}

	if request.Password == "" {
		log.Debugf("Password is empty")
		panic(exception.NewValidationError(messages.PasswordCannotBeEmpty))
		return
	}

	if request.Captcha == "" {
		log.Debugf("Captcha is empty")
		panic(exception.NewValidationError(messages.CaptchaCannotBeEmpty))
		return
	}

	if request.OID == "" {
		log.Debugf("OID is empty")
		panic(exception.NewValidationError(messages.OidCannotBeEmpty))
		return
	}

	// 校验验证码
	captcha, err := ctx.GetRedisClient().GetDel(ctx.GetRedisContext(), rdb.ConsoleLoginCaptcha.String(request.OID)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		panic(err)
		return
	}

	if captcha == "" || captcha != request.Captcha {
		panic(exception.NewValidationError(messages.CaptchaErrorOrDoesNotExist))
		return
	}
}

func (h *ConsoleHandler) Captcha(c *gin.Context) {
	var (
		err error
		log = logging.LoggerWithRequestID(c.Request.Context())
	)
	oid := ctx.GetOid(c)

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
	err = client.Set(ctx.GetRedisContext(), rdb.ConsoleLoginCaptcha.String(oid), answer, 5*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "captcha": b64s})
}

// Logout 后台管理系统登录
func (h *ConsoleHandler) Logout(c *gin.Context) {
	username := ctx.GetUsername(c)
	client := ctx.GetRedisClient()
	client.GetDel(ctx.GetRedisContext(), rdb.ConsoleToken.String(username))
}

// Current 后台管理系统登录
func (h *ConsoleHandler) Current(c *gin.Context) {
	var (
		username          = ctx.GetUsername(c)
		userRepo          = h.usersRepo
		usersProfilesRepo = h.userProfilesRepo
		err               error
		log               = logging.LoggerWithRequestID(c.Request.Context())
	)

	users, err := userRepo.SelectByUsername(username)
	if err != nil {
		log.Errorf(err.Error(), err)
		panic(exception.NewUnauthorizedError())
	}

	if users == nil {
		panic(exception.NewUnauthorizedError())
	}

	profiles, err := usersProfilesRepo.SelectByUserid(users.Id)
	if err != nil {
		panic(exception.NewUnauthorizedError())
	}

	if profiles == nil {
		panic(exception.NewUnauthorizedError())
	}

	c.JSON(200, dto.UsersDetail{
		Username:      users.Username,
		RecentLoginAt: users.LastLoginAt.Unix(),
		Email:         users.Email,
		SuperAdmin:    users.SuperAdmin,
		Profiles: dto.UsersProfiles{
			Birthday:     profiles.Birthday.Unix(),
			Gender:       profiles.Gender,
			Address:      profiles.Address,
			TOTPVerified: profiles.TotpVerified,
			TOTPEnabled:  profiles.TotpEnabled,
			Mobile:       profiles.Mobile,
			Nickname:     profiles.Nickname,
			Avatar:       profiles.Avatar,
			Bio:          profiles.Bio,
		},
	})
}

func (h *ConsoleHandler) GetTOTP(c *gin.Context) {
	var (
		usersRepo = h.usersRepo
		key       *otp.Key
		err       error
		log       = logging.LoggerWithRequestID(c.Request.Context())
	)

	username := ctx.GetUsername(c)
	log.Debugf("API '%s' username: '%s'", c.Request.URL.Path, username)

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

	if users.TotpSecret != "" {
		key, err = totp.Generate(totp.GenerateOpts{
			Issuer:      "RosenConsoleAdminPanel",
			AccountName: username,
			Secret:      []byte(users.TotpSecret),
		})
	} else {
		key, err = totp.Generate(totp.GenerateOpts{
			Issuer:      "RosenConsoleAdminPanel",
			AccountName: username,
		})
	}
	c.JSON(http.StatusOK, gin.H{
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
		if validate := totp.Validate(totpCode, users.TotpSecret); !validate {
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

	users.TotpSecret = secret

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

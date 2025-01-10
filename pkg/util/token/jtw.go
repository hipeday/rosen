package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var secret = []byte("rosen-jwt-secret") // JWT 密钥

func GenerateAdminPanelJWT(username, applicationName string, expiresAt time.Time) (string, error) {
	// 创建一个JWT的Claim
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    applicationName,                               // 颁发者
			Subject:   "Rosen admin panel console access token",      // 主题
			Audience:  jwt.ClaimStrings{"Rosen admin panel console"}, // 受众
			ExpiresAt: jwt.NewNumericDate(expiresAt),                 // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                // 签发时间
		},
	}

	// 创建一个JWT Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名并生成Token字符串
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT 解析 JWT
func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

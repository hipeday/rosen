package rdb

import "fmt"

type RedisKey string

const (
	ConsoleLoginCaptcha RedisKey = RedisKey("console:captcha:%s:login")
	ConsoleToken        RedisKey = RedisKey("console:%s:token")
)

func (k RedisKey) String(values ...any) string {
	return fmt.Sprintf(string(k), values...)
}

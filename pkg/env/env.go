package env

import (
	"os"
	"strings"
)

// GetOrDefault 获取环境变量如果环境变量 key 的内容 == "" 则返回 def
func GetOrDefault(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return def
}

// GetOrDefaultIgnoreCase 获取环境变量如果环境变量 key 的内容 == "" 则返回 def
func GetOrDefaultIgnoreCase(key string, def string) string {
	if value, exists := os.LookupEnv(strings.ToUpper(key)); exists {
		return value
	}
	return def
}

// Version 当前的Upay项目版本号
func Version() string {
	return os.Getenv("UPAY_VERSION")
}

// CurrentNodeId 当前节点ID 多节点唯一
func CurrentNodeId() int64 {
	return 0
}

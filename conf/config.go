package conf

import (
	"embed"
	"fmt"
	"github.com/hipeday/rosen/pkg/env"
	"gopkg.in/yaml.v3"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type DatabaseType string

const (
	PostgresSQL = "postgres"
	MySQL       = "mysql"
)

type Config struct {
	Application *Application `yaml:"application"`
	Database    *Database    `yaml:"database"`
	Server      *Server      `yaml:"server"`
	Logger      *Logging     `yaml:"logger"`
	Theme       *Theme       `yaml:"theme"`
	Redis       *Redis       `json:"redis"`
}

type Theme struct {
	Path    string `yaml:"path"`    // 主题路径
	Default string `yaml:"default"` // 默认主题名称
}

type Redis struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type Application struct {
	TimeZone string `yaml:"time_zone"`
	Version  string `yaml:"version"`
}

type Database struct {
	Type      DatabaseType  `yaml:"type"`
	Username  string        `yaml:"username"`
	Password  string        `yaml:"password"`
	Database  string        `yaml:"database"`
	Host      string        `yaml:"host"`
	Port      int16         `yaml:"port"`
	ParseTime bool          `yaml:"parse_time"`
	TimeZone  string        `yaml:"time_zone"`
	Pool      *DatabasePool `yaml:"pool"`
}

type DatabasePool struct {
	MaxConn     int `yaml:"max_conn"`      // 最大连接数
	MaxIdleConn int `yaml:"max_idle_conn"` // 最大空闲连接数
	MaxIdleTime int `yaml:"max_idle_time"` // 最大空闲时间 单位时 time.Hour
	MaxLife     int `yaml:"max_life"`      // 最大生命周期 单位秒 time.Second
}

type Server struct {
	IP   string `yaml:"ip"`
	Port int16  `yaml:"port"`
	Mode string `yaml:"mode"`
}

type Logging struct {
	// Encoding can be one "json" or "console". Defaults to "console"
	Encoding string `yaml:"encoding"`

	// Level configures the log level
	Level string `yaml:"level"`

	// Colors configures if color output should be enabled
	Colors *bool `yaml:"colors"`

	// time format
	TimeFormat string `yaml:"time_format"`
}

var (
	cfg  Config
	once sync.Once
	//go:embed *.yaml
	configFile embed.FS
)

func GetCfg() Config {
	once.Do(func() {
		loadConfigForEnv()
	})
	return cfg
}

// resolveEnvVars 替换 YAML 中的占位符 ${VAR:default} 为实际的环境变量值或默认值
func resolveEnvVars(data []byte) []byte {
	// 正则表达式匹配占位符 ${VAR:default}
	re := regexp.MustCompile(`\${(\w+):([^}]+)}`)

	// 替换逻辑
	processed := re.ReplaceAllFunc(data, func(match []byte) []byte {
		// 提取变量名和默认值
		matches := re.FindSubmatch(match)
		if len(matches) != 3 {
			return match // 无法解析的保留原样
		}

		envVar := string(matches[1])     // 环境变量名
		defaultVal := string(matches[2]) // 默认值

		// 获取环境变量值，如果不存在则使用默认值
		return []byte(strings.TrimLeft(env.GetOrDefault(envVar, defaultVal), " "))
	})

	return processed
}

// evaluateExpression evaluates a simple mathematical expression like "60*60*24".
func evaluateExpression(expr string) (int, error) {
	parts := strings.Split(expr, "*")
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid expression")
	}

	result := 1
	for _, part := range parts {
		value, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return 0, fmt.Errorf("invalid number in expression: %s", part)
		}
		result *= value
	}
	return result, nil
}

func loadConfigForEnv() {
	var (
		err error
	)
	data, err := configFile.ReadFile("config.yaml")

	// 解析 yaml 数据并处理占位符
	configData := resolveEnvVars(data)

	// 加载 YAML 数据
	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		panic(err)
	}

}

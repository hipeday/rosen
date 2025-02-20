package banner

import (
	"fmt"
	"github.com/hipeday/rosen/internal/ctx"
	"github.com/hipeday/rosen/pkg/env"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func readBannerFile() (string, error) {
	// 读取文件
	file, err := os.ReadFile("banner.txt")
	if err != nil {
		return "", err
	}
	return string(file), nil
}

// 使用反射获取嵌套字段的值
func getFieldValue(config interface{}, path []string) (string, error) {
	// 处理指针类型，确保解引用到结构体
	val := reflect.ValueOf(config)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // 解引用指针
	}

	// 逐层查找字段
	for _, p := range path {
		val = val.FieldByNameFunc(func(s string) bool {
			return strings.ToLower(s) == strings.ToLower(p) // 忽略大小写匹配字段名
		})
		if !val.IsValid() {
			return "", fmt.Errorf("field '%s' not found", p)
		}
	}
	return fmt.Sprintf("%v", val.Interface()), nil
}

// 处理占位符替换
func replacePlaceholders(input string) string {
	// 使用正则表达式匹配 ${...} 格式
	re := regexp.MustCompile(`\${(.*?)}`)
	input = re.ReplaceAllStringFunc(input, func(match string) string {
		// 提取占位符内容，例如：application.name
		placeholder := strings.Trim(match, "${}")
		parts := strings.Split(placeholder, ".")

		// 获取配置字段的值
		value, err := getFieldValue(ctx.GetConfig(), parts)
		if err != nil {
			// 如果找不到字段值，返回原占位符
			return match
		}
		return value
	})

	// 替换%...%格式的占位符为环境变量值
	reEnv := regexp.MustCompile(`%([^%]+)%`)
	return reEnv.ReplaceAllStringFunc(input, func(match string) string {
		// 获取环境变量的值
		envVar := strings.Trim(match, "%")
		envValue := env.GetOrDefaultIgnoreCase(envVar, match)

		if envValue == "" {
			// 如果环境变量为空，返回原占位符
			return match
		}
		return envValue
	})
}

func MakeBanner() {
	// 读取文件
	banner, err := readBannerFile()
	if err != nil {
		return
	}
	// 格式化
	banner = replacePlaceholders(banner)
	// 输出
	fmt.Println(text.FgHiGreen.Sprintf(banner))
}

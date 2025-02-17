package locales

import (
	"embed"
	"github.com/bytedance/sonic"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"strings"
)

// Language 语言类型
type Language string

var (
	EnUs Language = "en-US" // 英语(美国🇺🇸)
	ZhCn Language = "zh-CN" // 中文简体(中国🇨🇳)

	languages = []Language{EnUs, ZhCn}

	bundle *i18n.Bundle
	//go:embed *.json
	LocaleFS        embed.FS        // 扫描当前目录下面的所有json文件
	DefaultLanguage Language = EnUs // 默认使用英语
)

func (k Language) String() string {
	return string(k)
}

func init() {
	// 初始化 i18n Bundle，默认加载英文
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", sonic.Unmarshal)

	var (
		err error
	)

	// 加载语言文件
	_, err = loadFileFS(ZhCn) // 加载中文简体
	if err != nil {
		panic(err)
	}

	_, err = loadFileFS(EnUs) // 加载英文
	if err != nil {
		panic(err)
	}
}

func loadFileFS(lang Language) (*i18n.MessageFile, error) {
	return bundle.LoadMessageFileFS(LocaleFS, string(lang+".json"))
}

func Bundle() *i18n.Bundle {
	return bundle
}

func MatchLanguage(lang string) string {
	for _, item := range languages {
		if strings.ToLower(string(item)) == strings.ToLower(strings.ToLower(lang)) {
			return string(item)
		}
	}
	for _, item := range languages {
		if strings.Contains(strings.ToLower(string(item)), strings.ToLower(lang)) {
			return string(item)
		}
	}
	return string(DefaultLanguage)
}

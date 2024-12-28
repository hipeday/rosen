package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/pkg/util/file"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const rootPath = "/Users/coloxan/GolandProjects/hipeday/rosen"

func Run() {
	r := gin.Default()

	// 静态资源
	r.Static("/static", filepath.Join(rootPath, "themes", "__static__"))

	// 渲染 signin 页面，假设使用主题 'kuke'
	r.GET("/signin.html", func(c *gin.Context) {
		theme := c.DefaultQuery("theme", "kuke") // 这里可以从请求中或配置文件中动态获取当前主题
		// 构建对应主题下的页面路径
		// 加载模版路径
		r.LoadHTMLGlob(filepath.Join(rootPath, "themes", theme, "pages/*.html"))

		// 设置主题之后将主题对应静态资源复制一下
		err := file.CopyDir(filepath.Join(rootPath, "themes", theme, "static"), filepath.Join(rootPath, "themes", "__static__"))
		if err != nil {
			return
		}

		page := "signin.html"
		// 如果是移动端，使用移动端页面
		if isMobile(c.Request) {
			page = "signin-mobile.html"
		}

		// 渲染页面
		c.HTML(200, page, gin.H{
			"Theme": theme,
		})
	})

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// 检查是否是移动端
func isMobile(r *http.Request) bool {
	// 获取 User-Agent
	userAgent := r.Header.Get("User-Agent")

	// 定义常见的移动端关键字
	mobileKeywords := []string{"Mobile", "Android", "Silk/", "Kindle", "BlackBerry", "Opera Mini", "Opera Mobi", "iPhone", "iPad"}

	// 判断 User-Agent 是否包含移动端关键字
	for _, keyword := range mobileKeywords {
		if strings.Contains(userAgent, keyword) {
			return true
		}
	}
	return false
}

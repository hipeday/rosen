package banner

import (
	"fmt"
	"github.com/hipeday/rosen/conf"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strconv"
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

func bannerFormat(banner string) string {
	formattedBanner := strings.Replace(banner, "${application.version}", conf.GetCfg().Application.Version, -1)
	formattedBanner = strings.Replace(formattedBanner, "%pid%", strconv.Itoa(os.Getpid()), -1)
	return formattedBanner
}

func MakeBanner() {
	// 读取文件
	banner, err := readBannerFile()
	if err != nil {
		return
	}
	// 格式化
	banner = bannerFormat(banner)
	// 输出
	fmt.Println(text.FgHiGreen.Sprintf(banner))
}

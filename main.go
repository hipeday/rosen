package main

import (
	"embed"
	"github.com/hipeday/rosen/cmd/app"
	"github.com/hipeday/rosen/pkg/util/file"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	//go:embed *.env
	envFile embed.FS
)

func init() {
	// 检查 .env 文件是否存在
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		// 如果 .env 文件不存在，则从 example.env 复制一份
		if err := file.CopyFile("example.env", ".env"); err != nil {
			log.Fatalf("Failed to copy example.env to .env: %v", err)
		}
	}

	// 直接读取本地 .env 文件
	if err := loadEnv(".env"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
}

// 加载指定的环境文件
func loadEnv(fileName string) error {
	err := godotenv.Load(fileName)
	if err != nil {
		return err
	}
	log.Printf("Loaded environment variables from %s", fileName)
	return nil
}

func init() {
	data, err := envFile.ReadFile(".env")
	if err != nil || len(data) == 0 {
		data, err = envFile.ReadFile("example.env")
		if err != nil {
			log.Fatal("failed to read env file")
		}
	}
	// 将嵌入的 .env 数据写入到临时文件
	err = os.WriteFile(".env", data, 0644)
	if err != nil {
		log.Fatalf("failed to write release.env file: %v", err)
	}
	// 加载环境变量
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	app.Run()
}

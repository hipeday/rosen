package main

import (
	"github.com/hipeday/rosen/cmd/app"
	"github.com/hipeday/rosen/pkg/util/file"
	"github.com/joho/godotenv"
	"log"
	"os"
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

func main() {
	app.Run()
}

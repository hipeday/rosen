package main

import (
	"embed"
	"github.com/hipeday/rosen/cmd/app"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	//go:embed *.env
	envFile embed.FS
)

func init() {
	data, err := envFile.ReadFile(".env")
	if err != nil || len(data) == 0 {
		data, err = envFile.ReadFile("example.env")
		if err != nil {
			log.Fatal("failed to read env file")
		}
	}
	// 将嵌入的 .env 数据写入到临时文件
	err = os.WriteFile("release.env", data, 0644)
	if err != nil {
		log.Fatalf("failed to write release.env file: %v", err)
	}
	// 加载环境变量
	err = godotenv.Load("release.env")
	if err != nil {
		log.Fatalf("Error loading release.env file: %v", err)
	}
}

func main() {
	app.Run()
}

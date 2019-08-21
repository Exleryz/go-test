package main

import (
	"github.com/joho/godotenv"
	"go-test/server"
)

func main() {
	// 读取配置文文件
	godotenv.Load()
	r := server.Router()
	// 默认启动方式，包含 Logger、Recovery 中间件
	//r.Run(":8080")
	r.Run()
}

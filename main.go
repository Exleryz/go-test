package main

import (
	"go-test/server"
)

func main() {
	r := server.Router()
	// 默认启动方式，包含 Logger、Recovery 中间件
	//r.Run(":8080")
	r.Run()
}

package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	// base 字符串按照给定的进制进行解释
	// bitSize 表示的是整数取值范围，或者说整数的具体类型。取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64。
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})

	pong, err := client.Ping().Result()
	fmt.Printf("pong: %s, err: %v\n", pong, err)
	if err != nil {
		panic(err)
	}

	RedisClient = client
}

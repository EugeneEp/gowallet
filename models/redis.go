package models

import(
	"github.com/go-redis/redis"
	"os"
)


var RDB *redis.Client

func InitRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_ADDR"),
        Password: os.Getenv("REDIS_PASS"),
        DB:       0,
    })
}
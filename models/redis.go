package models

import(
	"github.com/go-redis/redis"
)


var RDB *redis.Client

func InitRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "foobared",
        DB:       0,
    })
}
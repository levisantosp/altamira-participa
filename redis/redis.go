package redis

import (
	"github.com/levisantosp/altamira-participa/utils"
	"github.com/redis/go-redis/v9"
)

var Client = redis.NewClient(&redis.Options{
	Addr: utils.Env.RedisAddr,
})

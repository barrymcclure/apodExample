package cache

import (
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

type Cacher struct {
	Address string
	Port    string
}

func NewCacher(address string, port string) *Cacher {
	c := Cacher{}
	c.init()
	return &c
}

func (c *Cacher) init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     c.Address + ":" + c.Port,
		Password: "",
		DB:       0,
	})
}

func (c *Cacher) FetchFromCache(key string, callback func() string) string {
	val, err := redisClient.Get(key).Result()
	if err = redis.Nil {
		val = callback()
	}
	return val
}

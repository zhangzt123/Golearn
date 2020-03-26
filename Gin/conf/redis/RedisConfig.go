package redis

import (
	"github.com/go-redis/redis/v7"
	c "github.com/zhangzt123/Golearn/Gin/conf"
	"log"
)

func NewRedisInstence() *redis.Client {
	addr := c.Conf.GetString("redis.addr")
	db := c.Conf.GetInt("redis.db")
	opt := redis.Options{}
	opt.Addr = addr
	opt.DB = db
	client := redis.NewClient(&opt)
	statecmd := client.Ping()
	err := statecmd.Err()
	if err != nil {
		panic(err)
	}
	defer func() {
		if e := recover(); e != nil {
			log.Fatal(e)
		}
	}()
	return client
}

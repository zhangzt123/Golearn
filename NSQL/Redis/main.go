package main

import (
	"fmt"

	redis "github.com/go-redis/redis/v7"        //redis
	msgpack "github.com/vmihailenco/msgpack/v4" //消息转换
)

var client *redis.Client

func init() {
	redisops := redis.Options{}
	redisops.Addr = "192.168.118.132:6379"
	redisops.DB = 0
	client = redis.NewClient(&redisops)
	str, err := client.Ping().Result()
	if err != nil {
		fmt.Printf(" %v\n %v \n", str, err)
	}

}

type users struct {
	username string
	Password string
}

func main() {

	user := users{username: "zhangsan", Password: "123"}

	bt, _ := msgpack.Marshal(user)
	stucmd := client.Set("key1", bt, 0)
	err := stucmd.Err()
	if err != nil {
		fmt.Println(err)
	}

	strcmd := client.Get("key1")
	key, e := strcmd.Result()
	if e == redis.Nil {
		fmt.Println(e)
	} else if e != nil {
		fmt.Println(e)
	} else {
		var str users
		e = msgpack.Unmarshal([]byte(key), &str)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(str)
	}
}

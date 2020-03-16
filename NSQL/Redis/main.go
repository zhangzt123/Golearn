/**go 连接redis
**/

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v7" //redis
	//消息转换
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

// type users struct {
// 	username string
// 	Password string
// }

func main() {

	//String

	//user := users{username: "zhangsan", Password: "123"}

	// bt, _ := msgpack.Marshal(user)
	/*
		bt, _ := msgpack.Marshal("user")
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
			// var str users
			var str string
			e = msgpack.Unmarshal([]byte(key), &str)
			if e != nil {
				fmt.Println(e)
			}
			fmt.Println(str)
		}
	*/
	//---------------------------------------------------------------------------------

	//hash
	/*
		//hashset
		intcmd := client.HSet("key", "user2", "wangwu", "user3", "zhouliu")
		i, hasherr := intcmd.Result()
		if hasherr == redis.Nil {
			fmt.Println(hasherr)
		} else if hasherr != nil {
			fmt.Println(hasherr)
		} else {
			// var str users
			if hasherr != nil {
				fmt.Println(hasherr)
			}
			fmt.Println(i)
		}
		//hashget

		slicecmd := client.HMGet("key", "user2", "user3")
		x, sliceerr := slicecmd.Result()
		if sliceerr != nil {
			fmt.Println(sliceerr)
		}
		s := make([]string, len(x))

		for i, str := range x {
			//借口类型需要类型断言
			if v, ok := str.(string); ok {
				s[i] = string(v)
			}
		}
		fmt.Println(s)

	*/

	//------------------------------------------------------------------------------------------

	//list

	// sw.Add(2)
	// go lp()
	// time.Sleep(1 * time.Second)
	// go rp()
	// sw.Wait()

	//------------------------------------------------------------------------------------------

	//set集和
	client.Set()
}

var stack string = "stack"

var sw sync.WaitGroup

//var c = make(chan string, 10)

//lpush
func lp() {
	defer sw.Done()
	for i := 0; i < 10; i++ {
		//c <- "1"
		intcmd := client.LPush(stack, i)
		j, err := intcmd.Result()
		if err == redis.Nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("向stack中推入%v 返回结果%v \n", i, j)

			time.Sleep(1 * time.Second)
		}

	}
	fmt.Printf("所有元素推入stack中 \n")

}

//rpop
func rp() {
	defer sw.Done()
forloop:
	for {
		//<-c
		stringcmd := client.RPop(stack)
		j, err := stringcmd.Result()
		if err == redis.Nil {
			fmt.Printf("打印redis错误：%v  stack中无元素 \n", err)
			time.Sleep(5 * time.Second)
			continue
		} else if err != nil {
			fmt.Println(err)
			break
		} else {
			switch {
			case strings.EqualFold(j, "49"):
				fmt.Printf("stack中弹出%v \n", j)
				fmt.Printf("stack中元素全部弹出 \n")
				break forloop
			default:
				fmt.Printf("stack中弹出%v \n", j)

			}

			time.Sleep(1 * time.Second)
		}

	}

}

/**
带缓存的管道

 main 方法会在大约20秒后连续打印10个数字之后会每10秒输出10个管道的内容
*/
package main

import (
	"fmt"
	"time"
)

// 声明一个带缓存的管道 大小100 管道缓存大于0 为异步（只有在缓存满了时候会发生阻塞）  0或者不填则为同步
// 可简写为 chan1 := make(chan int,100)
var chan1 chan int

func main() {
	chan1 = make(chan int, 100)
	go test1(chan1)
	time.Sleep(20 * 1e9)
	go test2(chan1)
	time.Sleep(100 * 1e9)
	fmt.Println("end")
}

// ch <-  表示从i流向管道
func test1(ch chan int) {
	for i := 1; i < 50; i++ {
		ch <- i
		time.Sleep(2 * 1e9)
	}
}

// i = <-ch 从管道中流向i
func test2(ch chan int) {
	var i int
	for {
		for j := 0; j < 10; j++ {
			i = <-ch
			//ch中没有数据 i会被阻塞无法从通道中获取 只能按照test1 的时间输出
			fmt.Println(i)
			//time.Sleep(1 * 1e9)
		}
		time.Sleep(10 * 1e9)
	}
}

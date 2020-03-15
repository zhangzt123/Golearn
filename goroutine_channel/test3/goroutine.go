/**
*通道阻塞
 */
package main

import (
	"fmt"
	"time"
)

// 声明一个管道
// 可简写为 chan1 := make(chan int)
var chan1 chan int

func main() {
	chan1 = make(chan int)
	go test1(chan1)
	go test2(chan1)
	time.Sleep(100 * 1e9)
	fmt.Println("end")
}

// ch <-  表示从i流向管道
func test1(ch chan int) {
	for i := 1; i < 1000; i++ {
		ch <- i
		time.Sleep(5 * 1e9)
	}
}

// i = <-ch 从管道中流向i
func test2(ch chan int) {
	var i int
	for {
		i = <-ch
		//ch中没有数据 i会被阻塞无法从通道中获取 只能按照test1 的时间输出
		fmt.Println(i)
		time.Sleep(1 * 1e9)
	}
}

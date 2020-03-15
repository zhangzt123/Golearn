/*
介绍管道
*/
package main

import (
	"fmt"
	"time"
)

// 声明一个管道
// 可简写为 chan1 := make(chan int)
// var chan1 chan int

func main() {
	chan1 := make(chan int)
	go test1(chan1)
	go test2(chan1)
	time.Sleep(10 * 1e9)
	close(chan1) //
	time.Sleep(10 * 1e9)
	fmt.Println("end")
}

// ch <-  表示从i流向管道
func test1(ch chan int) {

	defer func() {
		recover()
	}()
	for i := 0; i < 1000; i++ {
		ch <- i
		time.Sleep(1 * 1e9)
	}
}

// i = <-ch 从管道中流向i
func test2(ch chan int) {
	for {
		i, ok := <-ch
		fmt.Print(ok) //通道被关闭返回false
		if ok {
			fmt.Println(i)
		}
	}
}

/*
select{ case }  多个管道的处理
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
	chan2 := make(chan int)
	go test1(chan1)
	go test1(chan2)
	go test2(chan1, chan2)
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
func test2(ch, ch1 chan int) {
	for {
		select {
		case v, ok := <-ch:
			fmt.Printf("ch-%v-", ok) //通道被关闭返回false
			if ok {
				fmt.Println(v)
			}
		case v, ok := <-ch1:
			fmt.Printf("ch1-%v-", ok) //通道被关闭返回false
			if ok {
				fmt.Println(v)
			}
		default:
			fmt.Println("没有管道被处理")
		}
		time.Sleep(1 * 1e9)
	}
}

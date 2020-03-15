/*
sync.once
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 声明一个管道
// 可简写为 chan1 := make(chan int)
// var chan1 chan int

var wg sync.WaitGroup //等待组
var syncOnce sync.Once

func main() {
	runtime.GOMAXPROCS(2)
	chan1 := make(chan int)
	//chan2 := make(chan int)
	wg.Add(3)
	go test1(chan1, wg)
	go test1(chan1, wg)
	go test2(chan1, chan1, wg)
	wg.Wait() //wgadd为0时结束
	fmt.Println("end")
}

// ch <-  表示从i流向管道
func test1(ch chan int, wg sync.WaitGroup) {
	defer wg.Done() //wgadd减1
	defer func() {
		if err := recover(); err != nil {

			fmt.Printf("print err %v \n", err)
		}
	}()
	//defer syncOnce.Do(func() { close(ch) }) 当试用sync.once时保证关闭操作只会执行一次 不会发生panic

	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(0.1 * 1e9)
	}
	//close(ch) 多个线程同时关闭同一个管道是会发生panic 打印 print err send on closed channel
}

// i = <-ch 从管道中流向i
func test2(ch, ch1 chan int, wg sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1)
	defer ticker.Stop()
	for {
		select {
		case v, ok := <-ch:
			fmt.Printf("ch-%v-%v-", ok, <-ticker.C) //通道被关闭返回false
			if ok {
				fmt.Println(v)
			}
		case v, ok := <-ch1:
			fmt.Printf("ch1-%v-%v-", ok, <-ticker.C) //通道被关闭返回false
			if ok {
				fmt.Println(v)
			}
		default:
			fmt.Println("没有管道被处理")
		}
		time.Sleep(1 * 1e9)
	}
}

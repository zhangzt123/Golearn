/*
context利用函数内部的一个对外暴露的一个只允许接受的管道与其他线程间通信
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func loops(ctx context.Context) {

loop:
	for {
		fmt.Println(time.Now())
		time.Sleep(0.1 * 1e9)
		select {
		case <-ctx.Done():
			break loop
		default:
		}
	}
}

func main() {
	ctx, fun := context.WithCancel(context.Background())
	go loops(ctx)
	time.Sleep(10 * 1e9)
	fun()

}

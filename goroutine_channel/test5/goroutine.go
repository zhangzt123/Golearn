/**
带缓存的管道

 main 方法会在大约20秒后连续打印10个数字之后会每10秒输出一次管道的内容
*/
package main

import (
	"fmt"
	// "time"
)

// 声明一个带缓存的管道 大小100 管道缓存大于0 为异步（只有在缓存满了时候会发生阻塞）  0或者不填则为同步
// 可简写为 chan1 := make(chan int,100)
var chan1 chan int

func main() {
	// var i int
	slice1 := []byte("helo_world")
	chan1 = make(chan int, 100)
	go test1(slice1, chan1)
	if <-chan1 == 0 { //阻塞main等到test1完成后结束
		fmt.Println("end")
	}

}

// ch <-  表示从i流向管道
func test1(sl []byte, ch chan int) {
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%v", sl[i])
	}
	ch <- 0
}

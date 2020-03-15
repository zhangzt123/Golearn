/*介绍协程*/
package main

import (
	"fmt"
	"time"
)

//go 关键字启用一个go协程
func main() {
	go test1()
	go test2()
	time.Sleep(2220000000)
	fmt.Println("end")
}

func test1() {
	fmt.Println(1111111)
	time.Sleep(10)
	fmt.Println(2222222)
}

func test2() {
	fmt.Println(3333333)
	time.Sleep(10)
	fmt.Println(4444444)
}

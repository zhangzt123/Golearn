//互斥锁
package main

import "fmt"
import "sync"

var x int = 0

var wg sync.WaitGroup
var lock sync.Mutex // 锁

func add() {
	for i := 0; i < 1000000; i++ {
		lock.Lock() // 对共享资源加锁
		x++
		lock.Unlock() // 解锁
	}
	wg.Done()

}

func main() {

	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Print(x)
}

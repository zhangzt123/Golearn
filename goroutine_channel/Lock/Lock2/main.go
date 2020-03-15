//读写锁
package main

import "fmt"
import "sync"
import "time"

var x int = 0

var wg sync.WaitGroup
var lock sync.RWMutex // 锁

func add() {
	for i := 0; i < 100000; i++ {
		lock.Lock() // 对共享资源加锁
		x++
		time.Sleep(1 * 1e9)
		lock.Unlock() // 解锁
	}
	wg.Done()

}

func printt() {
	for i := 0; i < 200000; i++ {
		lock.RLock()
		fmt.Printf("%v ,", x)
		lock.RUnlock()
	}
	wg.Done()

}

func main() {

	wg.Add(3)
	go add()
	go add()
	go add()
	go printt()
	go printt()
	go printt()
	go printt()
	go printt()
	go printt()
	wg.Wait()

}

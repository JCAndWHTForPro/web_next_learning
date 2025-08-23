package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
    - 考察点 ： sync.Mutex 的使用、并发数据安全。
2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
    - 考察点 ：原子操作、并发数据安全。
*/

var lock sync.Mutex
var num = 0
var numNoLock int64 = 0

func noLockIncr(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i += 1 {
		atomic.AddInt64(&numNoLock, 1)
	}
}

func incr(wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	for i := 1; i <= 1000; i += 1 {
		num++
	}
}

func main() {
	var wg sync.WaitGroup
	//for i := 1; i <= 10; i += 1 {
	//	wg.Add(1)
	//	go incr(&wg)
	//}
	for i := 1; i <= 10; i += 1 {
		wg.Add(1)
		go noLockIncr(&wg)
	}
	wg.Wait()
	//fmt.Println("最终的结构是", num)
	fmt.Println("最终的noLock结构是", numNoLock)
}

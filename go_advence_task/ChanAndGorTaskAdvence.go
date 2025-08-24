package main

import (
	"fmt"
	"sync"
)

/**
题：并发输出1-100数字

要求一：启动2个协程，分别输出奇数和偶数(go_advence_task/GoroutineTask.go:15)

要求二：最多启动10个协程输出

要求三：最多启动10个协程顺序输出

*/
//要求一：启动2个协程，分别输出奇数和偶数
func cooroutine100EvenOddNum(isEven bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		even := (i%2 == 0)
		if even == isEven {
			fmt.Println(i)
		}
	}
}

// 要求二：最多启动10个协程输出
func cooroution100ConfuseNum(start int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start + 1; i <= start+10; i++ {
		fmt.Println(i)
	}

}

// 要求三：最多启动10个协程顺序输出
func cooroutineBych100OrderNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := <-ch
	for i := 0; i < 10; i += 1 {
		fmt.Println(start + i)
	}
	if start < 91 {
		ch <- start + 10
	}
}

func main() {
	var wg sync.WaitGroup
	//var ch = make(chan int)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go cooroutine100EvenOddNum(i%2 == 0, &wg)
	}
	//for i := 0; i < 10; i += 1 {
	//	wg.Add(1)
	//	//go cooroutineBych100OrderNum(ch, &wg)
	//	go cooroution100ConfuseNum(i*10, &wg)
	//}
	//ch <- 1
	wg.Wait()
	//close(ch)
}

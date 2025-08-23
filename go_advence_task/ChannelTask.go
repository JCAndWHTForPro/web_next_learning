package main

import (
	"fmt"
	"sync"
)

/**
1. 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
    - 考察点 ：通道的基本使用、协程间通信。
2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
    - 考察点 ：通道的缓冲机制。
*/

func sendInt(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 1; i <= 10; i += 1 {
		ch <- i
	}
}
func sendInt100(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 1; i <= 100; i += 1 {
		ch <- i
	}
}

func reciveInt(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for chInt := range ch {
		fmt.Println("是接收到了数据", chInt)
	}
}

func reciveInt100(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for chInt := range ch {
		fmt.Println("是接收到了数据", chInt)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go sendInt(ch, &wg)
	wg.Add(1)
	go reciveInt(ch, &wg)

	ch = make(chan int, 100)
	wg.Add(1)
	go sendInt100(ch, &wg)
	wg.Add(1)
	go reciveInt100(ch, &wg)

	wg.Wait()
}

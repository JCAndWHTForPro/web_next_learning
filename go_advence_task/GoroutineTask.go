package main

import (
	"fmt"
	"sync"
)

/**
1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
    - 考察点 ： go 关键字的使用、协程的并发执行。
2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
    - 考察点 ：协程原理、并发任务调度。
*/

func PrintNums(isEven bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 1 {
		if i%2 == 0 && isEven {
			fmt.Println("我是偶数线程的打印：", i)
		} else if i%2 == 1 && !isEven {
			fmt.Println("我是基数线程的打印：", i)
		}
	}
}

func taskExcutor(task []func(), wg *sync.WaitGroup) {
	for _, f := range task {
		go func() {
			defer wg.Done()
			f()
		}()
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go PrintNums(true, &wg)
	wg.Add(1)
	go PrintNums(false, &wg)

	wg.Add(3)
	taskExcutor([]func(){func() {
		fmt.Println("我是任务1")
	}, func() {
		fmt.Println("我是任务2")
	}, func() {
		fmt.Println("我是任务3")
	}}, &wg)
	wg.Wait()
	fmt.Println("所有线程都结束了")
}

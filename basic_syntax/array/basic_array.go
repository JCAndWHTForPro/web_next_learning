package main

import "fmt"

func main() {
	// 这种不是数组，是一个切片
	var arr = []int{}
	arr = append(arr, 123)
	fmt.Println(arr)

	// 这种是数组
	var arr_real = [3]int{123, 234, 345}
	fmt.Println(arr_real)

	// 这种是数组指针
	var arr_ptr *[3]int
	arr_ptr = &arr_real
	fmt.Println(*arr_ptr)

	// 这两种使用指针访问数组的方式是等价的，上面是语法糖
	arr_ptr[1] = 999
	(*arr_ptr)[1] = 101010
	fmt.Println(arr_ptr)

}

package main

import "fmt"

const max_size = 3

func main() {
	a := [3]int{123, 234, 345}
	var ptr_arr [max_size]*int
	//ptr_arr = &a
	for idx := 0; idx < len(a); idx += 1 {
		ptr_arr[idx] = &a[idx]
	}
	for idx := 0; idx < len(a); idx += 1 {
		fmt.Println(*ptr_arr[idx])
	}
}

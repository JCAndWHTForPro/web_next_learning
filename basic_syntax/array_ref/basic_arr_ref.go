package main

import "fmt"

func main() {

	// 这种
	var arr = [3]int{123, 234, 345}
	//arr = append(arr,345)
	fmt.Println(arr)
	modifyArr(arr)
	fmt.Println(arr)
	modifyArrRef(&arr)
	fmt.Println(arr)

}

func modifyArr(arr [3]int) {
	arr[1] = 23
}

func modifyArrRef(arr *[3]int) {
	(*arr)[1] = 22
}

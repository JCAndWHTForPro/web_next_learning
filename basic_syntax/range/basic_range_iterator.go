package main

import "fmt"

func main() {
	var slice = []int{123, 234, 345, 345}
	for idx, val := range slice {
		fmt.Printf("curr index:%d,value:%d\n", idx, val)
	}
	var arr = [4]string{"123", "asfd"}
	for idx, val := range arr {
		fmt.Printf("string iterator index:%d,value:%s\n", idx, val)
	}

	var str string = "jicheng"
	for i, c := range str {
		fmt.Printf("char iterator index:%d,char:%c\n", i, c)
	}
}

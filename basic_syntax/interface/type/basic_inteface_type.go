package main

import "fmt"

func printValueType(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

/*
*
空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。
1、任意类型都实现了空接口。
2、常用于需要存储任意类型数据的场景，如泛型容器、通用参数等。
*/
func main() {
	printValueType(11)
	printValueType("string")
	printValueType(12.234)
	printValueType([3]int{123, 234, 345})
	printValueType([]int{123, 234, 345})
}

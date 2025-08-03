package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}
	multiply := func(a, b int) int {
		return a * b
	}

	calc := func(oper func(a, b int) int, a, b int) int {
		return oper(a, b)
	}

	fmt.Printf("sum计算的结构是：%d\n", calc(add, 1, 2))
	fmt.Printf("multiply计算的结构是：%d\n", calc(multiply, 5, 2))
}

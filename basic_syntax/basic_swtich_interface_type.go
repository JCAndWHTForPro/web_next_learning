package main

import "fmt"

func main() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是：%T", i)
	case int:
		fmt.Printf("x的类型是：%s", i)
	}
}

package main

import "fmt"

/*
*
函数的定义
函数的值传递
函数的引用传递
*/
func main() {
	a := 100
	b := 200
	swap(a, b)
	fmt.Printf("a:%d,b:%d\n", a, b)
	swapp(&a, &b)
	fmt.Printf("a:%d,b:%d\n", a, b)
}

func swap(a int, b int) {
	a, b = b, a
}
func swapp(a *int, b *int) {
	*a, *b = *b, *a
}

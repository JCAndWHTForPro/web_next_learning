package main

import "fmt"

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	// 3<<2
	fmt.Println("k=", k)
	// 3<<3
	fmt.Println("l=", l)
}

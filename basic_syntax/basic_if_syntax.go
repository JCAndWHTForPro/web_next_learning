package main

import "fmt"

func main() {
	str := "456"
	if str == "123" {
		fmt.Println(true)
	} else if str == "456" {
		fmt.Println("false")
	} else {
		fmt.Println(123)
	}
}

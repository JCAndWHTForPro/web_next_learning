package main

import "fmt"

// 接口定义
type Animal interface {
	Yard()
}

type Dog struct {
	Sound string
}

func (d *Dog) Yard() {
	//d.Sound = "I am a dog"
	fmt.Println("I am" + d.Sound)
}

type Cat struct {
	Sound string
}

func (c *Cat) Yard() {
	//c.Sound = "I am a cat"
	fmt.Println("I am" + c.Sound)
}

func main() {
	var dog = Dog{Sound: " dog"}
	var annimal Animal = &dog
	annimal.Yard()
	cat := Cat{Sound: " cat"}
	annimal = &cat
	annimal.Yard()
}

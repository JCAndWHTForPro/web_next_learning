package main

import (
	"fmt"
	"math"
)

/**
1. 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
    - 考察点 ：接口的定义与实现、面向对象编程风格。
2. 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
    - 考察点 ：组合的使用、方法接收者。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	r float64
}

func (rec *Circle) Area() float64 {
	return rec.r * rec.r * math.Pi
}
func (rec *Circle) Perimeter() float64 {
	return 2 * rec.r * math.Pi
}

type Rectangle struct {
	length float64
	width  float64
}

func (rec *Rectangle) Area() float64 {
	return rec.width * rec.length
}
func (rec *Rectangle) Perimeter() float64 {
	return rec.width*2 + rec.length*2
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (per *Employee) PrintInfo() {
	name := per.Name
	age := per.Age
	id := per.EmployeeID
	fmt.Printf("我是一个员工，我的名字叫做%s，我年龄是%d，我的工号是%s", name, age, id)
}

func main() {
	shape := Shape(&Circle{
		r: 12.12,
	})
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
	shape = Shape(&Rectangle{
		length: 222.222,
		width:  111.222,
	})
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())

	var emp = Employee{
		Person: Person{
			Name: "123123",
			Age:  23,
		},
		EmployeeID: "123123123123123",
	}
	emp.PrintInfo()
}

/*变量的声明与初始化*/
package main

//包引入了也必须要被使用，否则会编译报错
import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

// 全局变量可以被声明但是不被使用
var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
	//局部变量声明了必须被使用
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, e, f, g, h)
}

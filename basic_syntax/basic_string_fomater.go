package main

import "fmt"

func main() {
	var num = 123
	var str = "hello world"
	var url = "http://www.baidu.com?code = %d & str = %s"
	var traget = fmt.Sprintf(url, num, str)
	fmt.Print(traget)
}

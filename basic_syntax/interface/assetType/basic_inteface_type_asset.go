package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{} = 123
	a, er := val.(string)
	if er {
		fmt.Println(a)
		fmt.Println(reflect.TypeOf(val))
	} else {
		fmt.Println("不是预想的类型")
		fmt.Println(reflect.TypeOf(val))
	}
}

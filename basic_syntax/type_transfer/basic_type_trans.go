package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数值类型的转换
	var a int = 10
	var b float64 = float64(a)
	fmt.Println(b)

	// 字符串转数字
	var stra string = "0123"
	if numa, ok := strconv.Atoi(stra); ok == nil {
		fmt.Println(numa)

	} else {
		fmt.Println("非数字类型")
	}

	var numa int = 123
	// 这种是数字转成字符串
	fmt.Println(strconv.Itoa(numa))
	// 这种是讲数字当成一个byte位来处理了，123对应的ascii码的字符是"{"
	fmt.Println(string(numa))

	str := "3.14"
	// 字符串转成具体的一个浮点数
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
	}

	nums := 3.14
	// 直接打印浮点数，展示是不带0 的
	fmt.Println(nums)
	// 浮点数格式化后转成字符串
	strs := strconv.FormatFloat(nums, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", nums, strs)

	// 接口类型的转换
	// 这里是直接将子类赋给了接口，因为Father这个结构体实现了Person的接口的方法，并且通过指针类型实现的
	var per Person = &Father{Say: "I am father"}
	speak, err := per.Speak("balabala")
	if err != nil {
		fmt.Println("异常：", err)
	} else {
		fmt.Println(speak)
	}
	// 这种是类似于java的向下转型，转型成具体子类，使用子类里面的具体的属性
	father, ok := per.(*Father)
	if !ok {
		fmt.Println("不是这个类型，你转换失败")
	} else {
		fmt.Println(father.Say)

	}

}

type Person interface {
	Speak(word string) (string, error)
}

type Father struct {
	Say string
}

func (fa *Father) Speak(word string) (string, error) {
	//panic("implement")
	s := fa.Say + ":" + word
	return s, nil
}

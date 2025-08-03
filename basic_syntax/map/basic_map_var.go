package main

import "fmt"

func main() {
	var siteMap = make(map[string]int)
	siteMap["fst"] = 1
	siteMap["sed"] = 2
	index, ok := siteMap["third"]
	if ok {
		fmt.Printf("当前取到了值：%d\n", index)
	} else {
		fmt.Printf("当前没取到了值\n")
	}
	_, isSedExit := siteMap["sed"]
	if isSedExit {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}

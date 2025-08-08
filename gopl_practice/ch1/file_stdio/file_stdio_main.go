package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
*
构建一个，识别输入的，如果是文件，则打开它，读出来，否则就是普通的
*/
func main() {
	maps := make(map[string]int)
	inpoutArgs := os.Args[1:]
	if len(inpoutArgs) == 0 {
		countLine(os.Stdin, maps)
	} else {
		for _, name := range inpoutArgs {
			f, errr := os.Open(name)
			if errr != nil {
				fmt.Fprintf(os.Stderr, "error,%v\n", errr)
				continue
			}
			countLine(f, maps)
		}
	}
	for k, v := range maps {
		fmt.Printf("当前的行是%s,统计到重复数量是：%d\n", k, v)
	}
}

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

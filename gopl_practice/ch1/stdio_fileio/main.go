package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lineCounter := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineCounter[scanner.Text()]++
	}
	for line, cnt := range lineCounter {
		fmt.Printf("当前行：%s,计数是：%d\n", line, cnt)
	}

}

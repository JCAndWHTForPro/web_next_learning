package main

import (
	"fmt"
	"math"
)

/*
*
给k个字符串，求出他们的最长公共前缀(LCP)
*/
func LongestCommonPrefix(strs []string) string {
	// write your code here
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var minLen = math.MaxInt
	for _, str := range strs {
		l := len(str)
		if l < minLen {
			minLen = l
		}
	}
	var start = 0
	for start < minLen {
		var eq = false
		var rst = strs[0]
		for index, str := range strs {
			if index == 0 {
				continue
			}
			eq = str[start] == rst[start]
			if !eq {
				break
			}
		}
		if !eq {
			break
		}
		start += 1
	}

	return strs[0][0:start]
}

func main() {
	fmt.Println(LongestCommonPrefix([]string{"ABCDEFG", "ABCEFG", "ABCEFA"}))
}

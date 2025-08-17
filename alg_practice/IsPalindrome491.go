package main

import (
	"fmt"
	"strconv"
)

/*
*
判断一个正整数是不是回文数。

回文数的定义是，将这个数反转之后，得到的数仍然是同一个数。
*/
func IsPalindrome(num int) bool {
	// write your code here
	itoa := strconv.Itoa(num)
	l := len(itoa)
	start, end := 0, l-1
	for start < end {
		if itoa[start] != itoa[end] {
			return false
		}
		start += 1
		end += 1
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(11))
}

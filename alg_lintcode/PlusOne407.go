package main

import "fmt"

/*
描述
给定一个非负数，表示一个数字数组，在该数的基础上+1，返回一个新的数组。

该数字按照数位高低进行排列，最高位的数在列表的最前面。
*/
func PlusOne(digits []int) []int {
	// write your code here
	l := len(digits)
	if l == 0 {
		return append([]int{}, 1)
	}
	var hasTen = false
	newNum := digits[l-1] + 1
	if newNum >= 10 {
		hasTen = true
		newNum = 0
	}
	digits[l-1] = newNum
	for i := l - 2; hasTen && i >= 0; i -= 1 {
		lastDg := digits[i] + 1
		if lastDg < 10 {
			hasTen = false
			digits[i] = lastDg
		} else {
			digits[i] = 0
		}
	}
	if hasTen {
		digits = append([]int{1}, digits...)
	}
	return digits
}
func main() {
	fmt.Println(PlusOne([]int{2, 9, 9}))
}

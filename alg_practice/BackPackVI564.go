package main

import "fmt"

/*
*描述
给出一个都是正整数的数组 nums，其中没有重复的数。从中找出所有的和为 target 的组合个数。

一个数可以在组合中出现多次。
数的顺序不同则会被认为是不同的组合。
1≤nums.length≤100
1≤nums[i]≤1000
1≤target≤1000
*/
func BackPackVI(nums []int, target int) int {
	// write your code here
	n := len(nums)
	if n == 0 {
		return -1
	}
	m := target
	dp := make([]int, m+1)
	dp[0] = 1
	for j := 1; j <= m; j += 1 {
		for _, num := range nums {
			if j >= num {
				dp[j] += dp[j-num]
			}
		}
	}
	return dp[m]
}
func main() {
	fmt.Println(BackPackVI([]int{1, 2, 4}, 4))
}

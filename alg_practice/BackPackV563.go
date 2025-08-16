package main

import "fmt"

/*
*描述
给出 n 个物品, 以及一个数组, nums[i] 代表第i个物品的大小,
保证大小均为正数, 正整数 target 表示背包的大小, 找到能填满背包的方案数。
每一个物品只能使用一次

1≤nums.length≤1000
1≤target≤1000
*/
func BackPackV(nums []int, target int) int {
	// write your code here
	n := len(nums)
	if n == 0 {
		return -1
	}
	m := target
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i += 1 {
		dp[i%2][0] = 1
	}
	for i := 1; i <= n; i += 1 {
		for j := 0; j <= m; j += 1 {
			dp[i%2][j] = dp[(i-1)%2][j]
			if j >= nums[i-1] {
				dp[i%2][j] += dp[(i-1)%2][j-nums[i-1]]
			}
		}
	}
	return dp[n%2][m]
}

func main() {
	fmt.Println(BackPackV([]int{1, 2, 3, 3, 7}, 7))
}

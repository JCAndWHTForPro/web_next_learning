package main

/*
给一 只含有正整数 的 非空 数组, 找到这个数组是否可以划分为 两个 元素和相等的子集。

所有数组元素不超过100.
数组大小不超过200.
*/
func CanPartition(nums []int) bool {
	// write your code here
	var sum int = 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	m := sum / 2
	n := len(nums)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	for i := 0; i <= n; i += 1 {
		dp[i][0] = true
	}

	for i := 1; i <= n; i += 1 {
		for j := 1; j <= m; j += 1 {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][m]
}

package main

/*
*描述
给定 n 个不同的正整数，整数 k(k≤n)以及一个目标数字 target。
在这 n 个数里面找出 k 个数，使得这 k 个数的和等于目标数字，求问有多少种方案？
*/
func KSum(a []int, k int, target int) int {
	// write your code here
	l := len(a)
	if l == 0 {
		return 0
	}
	dp := make([][][]int, l+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, target+1)
		}
	}

	for i := 0; i <= l; i += 1 {
		dp[i][0][0] = 1
	}
	for i := 1; i <= l; i += 1 {
		for j := 1; j <= k; j += 1 {
			for t := 1; t <= target; t += 1 {
				dp[i][j][t] = dp[i-1][j][t]
				if t >= a[i-1] {
					dp[i][j][t] += dp[i-1][j-1][t-a[i-1]]
				}
			}
		}
	}

	return dp[l][k][target]
}

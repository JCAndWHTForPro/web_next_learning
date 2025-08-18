package main

func BackPackIII(a []int, v []int, m int) int {
	// write your code here
	n := len(a)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	//for j := 0; j <= m; j += 1 {
	//	dp[0][j] = 0
	//}
	for i := 1; i <= n; i += 1 {
		for j := 1; j <= m; j += 1 {
			dp[i][j] = dp[i-1][j]
			if j >= a[i-1] && dp[i][j-a[i-1]]+v[i-1] > dp[i][j] {
				dp[i][j] = dp[i][j-a[i-1]] + v[i-1]
			}
		}
	}

	return dp[n][m]
}

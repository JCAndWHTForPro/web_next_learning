package main

/*
*描述
在 n 个物品中挑选若干物品装入背包，最多能装多满？假设背包的大小为m，每个物品的大小为Ai
（每个物品只能选择一次且物品大小均为正整数）

你不可以将物品进行切割
n<1000
m<1e5

当前做法是选择前i个数能否组转出j这么大的背包（精准）
*/
func BackPack(m int, a []int) int {
	// write your code here
	if len(a) == 0 {
		return 0
	}
	n := len(a)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i += 1 {
		for j := 0; j <= m; j += 1 {
			if j >= a[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-a[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	for i := m; i >= 0; i -= 1 {
		if dp[n][i] {
			return i
		}
	}
	return -1
}

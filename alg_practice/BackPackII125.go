package main

/*
*描述
有 n 个物品和一个大小为 m 的背包. 给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值.

问最多能装入背包的总价值是多大?

A[i], V[i], n, m 均为整数
你不能将物品进行切分
你所挑选的要装入背包的物品的总大小不能超过 m
每个物品只能取一次
m<=1000\
len(A),len(V)<=100
*/
func BackPackII(m int, a []int, v []int) int {
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
			if j >= a[i-1] && dp[i-1][j-a[i-1]]+v[i-1] > dp[i][j] {
				dp[i][j] = dp[i-1][j-a[i-1]] + v[i-1]
			}
		}
	}

	return dp[n][m]
}

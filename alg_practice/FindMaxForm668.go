package main

/*
*描述
在计算机世界中, 由于资源限制, 我们一直想要追求的是产生最大的利益.
现在，假设你分别是 m个 0 和 n个 1 的统治者. 另一方面, 有一个只包含 0 和 1 的字符串构成的数组.
现在你的任务是找到可以由 m个 0 和 n个 1 构成的字符串的最大个数. 每一个 0 和 1 均只能使用一次

给出的 0 和 1 的个数不会超过 100
给出的字符串数组的大小不会超过 600
*/
func FindMaxForm(strs []string, m int, n int) int {
	// write your code here
	l := len(strs)
	if l == 0 {
		return 0
	}
	dp := make([][][]int, l+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 1; i <= l; i += 1 {
		ints := count(strs[i-1])
		for j := 0; j <= m; j += 1 {
			for k := 0; k <= n; k += 1 {
				dp[i][j][k] = dp[i-1][j][k]
				if j >= ints[0] && k >= ints[1] && dp[i-1][j-ints[0]][k-ints[1]]+1 > dp[i][j][k] {
					dp[i][j][k] = dp[i-1][j-ints[0]][k-ints[1]] + 1
				}
			}
		}
	}
	return dp[l][m][n]
}

func count(str string) []int {
	rst := make([]int, 2)
	for _, ch := range str {
		i := ch - '0'
		rst[i]++
	}
	return rst
}

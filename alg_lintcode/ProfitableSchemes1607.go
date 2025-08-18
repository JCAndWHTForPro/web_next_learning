package main

/*
*描述
有 G 名黑帮成员要计划接下来的犯罪活动。
给出两个等长的数组groups 和 profit，它们的含义为第 i 个犯罪需要投入groups[i]个成员，来获得profit[i]的盈利。
一名成员只能同时参加一场犯罪活动。

黑帮的目的是至少获得 P 的盈利，你的任务是计算有多少种方案可以选择。
因为答案很大，所以答案要对 1000000007 取模。

1 <= G <= 100
0 <= P <= 100
1 <= group[i] <= 100
0 <= profit[i] <= 100
1 <= group.length = profit.length <= 100
*/
func ProfitableSchemes(g int, p int, group []int, profit []int) int {
	// Write your code here.
	// write your code here
	l := len(group)
	if l == 0 {
		return 0
	}
	var mod = 1000000007
	dp := make([][][]int, l+1)
	for i := range dp {
		dp[i] = make([][]int, g+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, p+1)
		}
	}
	dp[0][0][0] = 1

	for i := 1; i <= l; i += 1 {
		person := group[i-1]
		prof := profit[i-1]
		for j := 0; j <= g; j += 1 {
			for t := 0; t <= p; t += 1 {
				dp[i][j][t] = dp[i-1][j][t]
				if j >= person {
					r := t - prof
					if r < 0 {
						r = 0
					}
					dp[i][j][t] = (dp[i][j][t] + dp[i-1][j-person][r]) % mod
				}
			}
		}
	}

	var sum = 0
	for i := 0; i <= g; i += 1 {
		sum = (sum + dp[l][i][p]) % mod
	}
	return sum
}

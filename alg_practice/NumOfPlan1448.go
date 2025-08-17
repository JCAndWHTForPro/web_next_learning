package main

/*
*
现在有一个卡牌游戏，先给出卡牌的数量n，再给你两个非负整数totalProfit、totalCost，
然后给出每张卡牌的利润值 a[i]和成本值b[i]，现在可以从这些卡牌中任意选择若干张牌，组成一个方案，
问有多少个方案满足所有选择的卡牌利润和大于totalProfit且成本和小于totalCost。
由于这个数可能很大，因此只需返回方案数对1e9 + 7取模的结果。

0≤n≤100
0≤totalProfit≤100
0≤totalCost≤100
0≤a[i]≤100
0≤b[i]≤100
*/
const MOD = 1000000007

func NumOfPlan(n int, totalProfit int, totalCost int, a []int, b []int) int {
	// Write your code here
	totalProfit++
	var dp [2][110][110]int
	for i := 0; i <= totalCost; i += 1 {
		dp[0][0][i] = 1
	}
	for i := 1; i <= n; i += 1 {
		for j := 0; j <= totalProfit; j += 1 {
			for t := 0; t <= totalCost; t += 1 {
				dp[i%2][j][t] = dp[(i-1)%2][j][t]
				if t >= b[i-1] {
					jj := j - a[i-1]
					if jj < 0 {
						jj = 0
					}
					dp[i%2][j][t] = (dp[i%2][j][t] + dp[(i-1)%2][jj][t-b[i-1]]) % MOD
				}

			}
		}
	}
	return dp[n%2][totalProfit][totalCost-1]
}

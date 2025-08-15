package main

import (
	"fmt"
	"math"
)

/*
*
描述
给出不同面额的硬币以及一个总金额. 写一个方法来计算给出的总金额可以换取的最少的硬币数量. 如果已有硬币的任意组合均无法与总金额面额相等, 那么返回 -1.

你可以假设每种硬币均有无数个
总金额不会超过50000
硬币的种类数不会超过500, 每种硬币的面额不会超过100
*/
func CoinChange(coins []int, amount int) int {
	// write your code here

	n := len(coins)
	if n == 0 {
		return -1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= n; i += 1 {
		for j := 0; j <= amount; j += 1 {
			if j == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt
			}
		}
	}
	for i := 1; i <= n; i += 1 {
		for j := 1; j <= amount; j += 1 {
			dp[i][j] = dp[i-1][j]
			if j >= coins[i-1] && dp[i][j-coins[i-1]] != math.MaxInt {
				dp[i][j] = min(dp[i][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}

	if dp[n][amount] == math.MaxInt {
		return -1
	}
	return dp[n][amount]
}

func main() {
	fmt.Println(CoinChange([]int{1, 9}, 0))
}

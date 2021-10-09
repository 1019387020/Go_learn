package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, 3)
	for k := 0; k <= 2; k++ {
		dp[k] = make([]int, 2)
	}
	dp[1][0] = 0
	dp[1][1] = -prices[0]
	dp[2][0] = 0
	dp[2][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[2][0] = max(dp[2][0], dp[2][1]+prices[i])
		dp[2][1] = max(dp[2][1], dp[1][0]-prices[i])
		dp[1][0] = max(dp[1][0], dp[1][1]+prices[i])
		dp[1][1] = max(dp[1][1], dp[0][0]-prices[i])
	}
	return dp[2][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	prices := []int{1}
	res := maxProfit(prices)
	fmt.Println(res)
}
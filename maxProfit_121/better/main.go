package main

import "fmt"

//时间复杂度：O(n)
//空间复杂度：O(n)
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	prices := []int{7,1,5,3,6,4}
	res := maxProfit(prices)
	fmt.Println(res)
}
package main

import "fmt"

func maxProfit(k int, prices []int) int {
	if prices == nil {
		return 0
	}
	length := len(prices)
	if k >= length/2 {
		return _maxProfit(prices)
	}
	dp := make([][][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 1; i <= k; i++ {
		dp[0][i][0] = 0
		dp[0][i][1] = -prices[0]
	}
	for i := 1; i < length; i++ {
		for j := k; j > 0; j-- {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[length-1][k][0]
}

func _maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	length := len(prices)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[length-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	k := 2
	prices := []int{3,2,6,5,0,3}
	res := maxProfit(k, prices)
	fmt.Println(res)
}

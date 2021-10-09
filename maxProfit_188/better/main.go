package main

import "fmt"

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	length := len(prices)
	if k > length/2 {
		return _maxProfit(prices)
	}
	dp := make([][]int, k+1)
	for j := 0; j <= k; j++ {
		dp[j] = make([]int, 2)
	}
	for j := 1; j <= k; j++ {
		dp[j][0] = 0
		dp[j][1] = -prices[0]
	}
	for i := 1; i < length; i++ {
		for j := k; j > 0; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}
	return dp[k][0]
}

func _maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	length := len(prices)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]+prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]-prices[i])
	}
	return dp[(length-1)%2][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	k := 2
	prices := []int{2,4,1}
	res := maxProfit(k, prices)
	fmt.Println(res)
}
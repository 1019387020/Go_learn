package main

import "fmt"

//不限次数的买进卖出
//与只记录一次买进卖出的最大区别，在动态规划状态方程的变化
//记录一次：dp[i][1] = max(dp[i-1][1], -prices[i])
//记录多次：dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]+prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]-prices[i])
	}
	return dp[(n-1)%2][0]
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
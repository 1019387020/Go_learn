package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount+1
	}
	dp[0] = 0

	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	coins := []int{2}
	amount := 3
	res := coinChange(coins, amount)
	fmt.Println(res)
}

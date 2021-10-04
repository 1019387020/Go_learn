package main

import "fmt"

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		// if i % 7 == 0 {
		// 	dp[i] = 0
		// 	continue
		// }
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(7))
}
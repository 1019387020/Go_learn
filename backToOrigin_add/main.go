package main

import "fmt"

// m 为圆环中点的个数， n 为需要走的步数
func backToOrigin(m int, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1
	for i := 1; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			dp[i][j] = dp[i-1][(j-1+m)%m] + dp[i-1][(j+1+m)%m]
		}
	}
	fmt.Println(dp)
	return dp[n][0]
}

func main() {
	res := backToOrigin(10, 2)
	fmt.Println(res)
}
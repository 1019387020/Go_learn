package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = math.MaxInt32
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = math.MaxInt32
	}

	dp[1][1] = grid[0][0]
	for i := 1; i <= m; i++ {
		for j := 1; j <=n; j++ {
			if i == 1 && j ==1 {
				continue
			}
			dp[i][j] = grid[i-1][j-1] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m][n]
}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	// grid := [][]int{
	// 	{1,3,1},
	// 	{1,5,1},
	// 	{4,2,1},
	// }
	grid := [][]int{
		{1,2,3},
		{4,5,6},
	}
	res := minPathSum(grid)
	fmt.Println(res)
}
package main

import "fmt"

// 原地 DP
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0] 
	}
	for j := 1; j < n; j++ {
		//grid[0][j] += grid[0][j] 错误：j
		grid[0][j] += grid[0][j-1]
	}

	for	i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	fmt.Println(grid)
	return grid[m-1][n-1]
}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	grid := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	// grid := [][]int{
	// 	{1,2,3},
	// 	{4,5,6},
	// }
	res := minPathSum(grid)
	fmt.Println(res)
}
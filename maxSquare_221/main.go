package main

import (
	"fmt"
)

func maxSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	res := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if matrix[i][n-1] == '1' {
			dp[i][n-1] = 1
		}
		if dp[i][n-1] > res {
			res = dp[i][n-1]
		}
	}
	for j := 0; j < n; j++ {
		if matrix[m-1][j] == '1' {
			dp[m-1][j] = 1
		}
		if dp[m-1][j] > res {
			res = dp[m-1][j]
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			}
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(
					dp[i+1][j],
					dp[i][j+1],
					dp[i+1][j+1],
				)
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	fmt.Println(dp)
	return res * res
}

func min(x, y, z int) int {
	return minTwo(x, minTwo(y, z))
}

func minTwo(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	// matrix := [][]byte{
	// 	{'0', '1'},
	// 	{'1', '0'},
	// }
	res := maxSquare(matrix)
	fmt.Println(res)
}

package main

import "fmt"

// 原地dp , 优化空间复杂度为 O(1)
func maxSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	res := 0
	for i := 0; i < m; i++ {
		if matrix[i][n-1] == '1' {
			res = 1
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[m-1][j] == '1' {
			res = 1
			break
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if matrix[i][j] == '0' {
				continue
			}
			if matrix[i][j] == '1' {
				matrix[i][j] = 1 + min(
					matrix[i+1][j],
					matrix[i][j+1],
					matrix[i+1][j+1],
				)
			}
			if int(matrix[i][j]-'0') > res {
				res = int(matrix[i][j] - '0')
			}
		}
	}
	fmt.Printf("%c\n", matrix)
	return res * res
}

func min(x, y, z byte) byte {
	return minTwo(x, minTwo(y, z))
}

func minTwo(x, y byte) byte {
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

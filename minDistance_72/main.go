package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j]+1,   // 删除
					dp[i][j-1]+1,   // 插入
					dp[i-1][j-1]+1, // 替换
				)
			}
		}
	}

	return dp[m][n]
}

func min(x, y, z int) int {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}

func main() {
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistance(word1, word2))
}
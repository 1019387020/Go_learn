package main

import "fmt"

func longestCommonSubsquence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 数组索引从 0 开始，因此不是 text1[i] = text2[j]
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	text1, text2 := "abc", "abc" 
	res := longestCommonSubsquence(text1, text2)
	fmt.Println(res)
}
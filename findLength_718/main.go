package main

import "fmt"

// 动态规划解法
func findLength(nums1 []int, nums2 []int) (int, int) {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	start, res := 0, 0
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > res {
				start = i
				res = dp[i][j]
			}
		}
	}
	
	return res, start
}

func main() {
	nums1 := []int{1,2,3,2,1}
	nums2 := []int{3,2,1,4,7}
	res, start := findLength(nums1, nums2)
	fmt.Println(res, nums1[start : start+res])
}
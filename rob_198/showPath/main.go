package main

import (
	"fmt"
	"sort"
)

func rob(nums []int) int {
	m := len(nums)
	if m == 0 {
		return 0
	}
	if m == 1 {
		return nums[0]
	}
	if m == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = nums[0]
	dp[1][0] = nums[0]
	dp[1][1] = nums[1]
	for i := 2; i < m; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	res := max(dp[m-1][0], dp[m-1][1])
	target := res
	
	path := make([]int, 0)
	for i := m-1; i >= 0; i-- {
		if (dp[i][1] == target) {
			path = append(path, i)
			target -= nums[i]
		}
	}
	// fmt.Println(nums)
	// fmt.Println(dp)
	sort.Ints(path)
	// fmt.Println(path)
	// fmt.Println()
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2,1,1,2}
	res := rob(nums)
	fmt.Println(res)
}

package main

import (
	"fmt"
)

func rob(nums []int) int {
	m := len(nums)
	if m == 0 {
		return 0
	}
	if m == 1 {
		return nums[0]
	}
	dp := make([]int, m)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < m; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	// fmt.Println(nums)
	// fmt.Println(dp)
	return dp[m-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2,1,1,2}
	// nums := []int{2,1}
	res := rob(nums)
	fmt.Println(res)
}

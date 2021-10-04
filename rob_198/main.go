package main

import (
	"fmt"
)

func rob(nums []int) int {
	m := len(nums)
	dp := make([]int, m)
	if m < 2 {
        return nums[m-1]
    }
	dp[0], dp[1] = nums[0], nums[1]
	res := max(dp[0], dp[1])
	for i := 2; i < m; i++ {
		for j := 0; j <= i-2; j++ {
			dp[i] = max(dp[i], nums[i]+dp[j])
		}
		res = max(res, dp[i])
	}
	fmt.Println(nums)
	fmt.Println(dp)
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

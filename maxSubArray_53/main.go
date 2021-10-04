package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	res := math.MinInt64
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}
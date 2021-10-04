package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) (int, []int) {
	n := len(nums)
	if n == 0 {
		return 0, nums
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	left := 0
	for i := 1; i < n; i++ {
		//dp[i] = max(nums[i], dp[i-1]+nums[i])
		if nums[i] > dp[i-1]+nums[i] {
			dp[i] = nums[i]
			left = i
		} else {
			dp[i] = dp[i-1]+nums[i]
		}

	}

	res := math.MinInt64
	right := 0
	for i := 0; i < n; i++ {
		// res = max(res, dp[i])
		if dp[i] > res {
			res = dp[i]
			right = i
		}
		continue
	}
	return res, nums[left : right+1]
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}
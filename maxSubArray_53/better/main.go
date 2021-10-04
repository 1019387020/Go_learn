package main

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp0, dp1, res := nums[0], 0, nums[0]
	for i := 1; i < len(nums); i++ {
		dp1 = max(nums[i], nums[i]+dp0)
		dp0 = dp1
		res = max(res, dp1)
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
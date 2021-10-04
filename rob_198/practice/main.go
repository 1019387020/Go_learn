package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	track := make([]bool, len(nums))
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		if dp[i-2]+nums[i] > dp[i-1] {
			track[i] = true
		}
	}
	if track[2] {
		track[0] = true
	}
	fmt.Println(nums)
	fmt.Println(dp)
	fmt.Println(track)
	return dp[len(nums)-1]
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
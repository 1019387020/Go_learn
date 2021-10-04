package main

import "fmt"

// 时间复杂度：O(n^2)，空间复杂度 O(n)
func longthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)), 0
	// for i := 0; i < len(nums); i++ {
	// 	dp[i] = 1
	// }

	for i := 0; i < len(nums); i++ {
		// 初始化 dp[i] , 效果和外面的初始化方法相同
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(longthOfLIS(nums))
}
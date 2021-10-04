package main

import (
	"fmt"
	"sort"
)

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]
	dp[1][0] = nums[0]
	dp[1][1] = nums[1]
	for i := 2; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	res := max(dp[n-1][0], dp[n-1][1])
	target := res
	path := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		if dp[i][1] == target {
			path = append(path, i)
			target -= nums[i]
		}
	}
	sort.Ints(path)
	fmt.Println(path)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 1, 1, 2}
	res := rob(nums)
	fmt.Println(res)
}
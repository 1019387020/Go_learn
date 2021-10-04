package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	left, sum, res := 0, 0, len(nums)+1
	for right, v := range nums {
		sum += v
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	target := 11
	nums := []int{1, 2, 3, 4, 5}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}

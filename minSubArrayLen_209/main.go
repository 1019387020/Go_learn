package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	sum, res := 0, len(nums)+1
	left, right, windowLen := 0, 0, 0
	for right < len(nums) {
		if sum < target {
			sum += nums[right]
			right++
			windowLen = right - left
			if sum >= target && windowLen < res {
				res = windowLen
			}
		}

		for sum >= target {
			sum -= nums[left]
			left++
			windowLen = right - left
			if sum >= target && windowLen < res {
				res = windowLen
			}
		}
	}

	if res == len(nums)+1 {
		return 0
	} else {
		return res
	}
}

func main() {
	target := 11
	nums := []int{1, 2, 3, 4, 5}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}

package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	sz := len(nums)
	res := make([]int, 0)
	if sz < 2 {
		return res
	}
	lo, hi := 0, sz-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum < target {
			lo++
		} else if sum > target {
			hi--
		} else {
			res = append(res, lo, hi)
			break
		}
	}
	return res
}

func main() {
	nums := []int{3, 2, 4}
	res := twoSum(nums, 6)
	fmt.Println(res)
}
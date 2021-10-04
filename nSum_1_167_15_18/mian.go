package main

import (
	"fmt"
	"sort"
)


func nSumTarget(nums []int, n int, start int, target int) [][]int {
	sz := len(nums)
	res := make([][]int, 0)
	if n < 2 || sz < n {
		return res
	}
	if n == 2 {
		lo, hi := start, sz-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				res = append(res, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 4, 0, target)
}

func main() {
	//nums := []int{-1,0,1,2,-1,-4}
	// res := threeSum(nums)
	
	//nums := []int{1,0,-1,0,-2,2}
	//res := fourSum(nums, 0)
	nums := []int{2,2,2,2,2}
	res := fourSum(nums, 8)
	
	fmt.Println(res)
}
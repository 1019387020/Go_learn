package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res, track := [][]int{}, []int{}
	sort.Ints(candidates)
	backtrack(candidates, target, 0, track, &res)
	return res
}

func backtrack(nums []int, target int, index int, track []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			temp := make([]int, len(track))
			copy(temp, track)
			*res = append(*res, temp)
		}
		return
	}

	for i := index; i < len(nums); i++ {
		if nums[i] > target {
			continue
		}
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backtrack(nums, target-nums[i], i+1, track, res)
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(nums, target))
}

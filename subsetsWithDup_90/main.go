package main

import (
	"sort"
	"fmt"
)

func subsetsWithDup(nums []int) [][]int {
	track, res := []int{}, [][]int{}
	sort.Ints(nums)
	backtrack(nums, 0, track, &res)
	return res
}

func backtrack(nums []int, index int, track []int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backtrack(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{1,2,2}
	fmt.Println(subsetsWithDup(nums))
}
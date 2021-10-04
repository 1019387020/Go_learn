package main

import (
	"fmt"	
	"sort"
)


func permuteUnique(nums []int) [][]int {
	used, res, track := make([]bool, len(nums)), [][]int{}, []int{}
	sort.Ints(nums)
	backtrack(nums, track, &res, &used)
	return res
}

func backtrack(nums []int, track []int, res *[][]int, used *[]bool) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		//fmt.Printf("track:%d\n", temp)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 排除不合法选择
		if (*used)[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
			continue
		}

		track = append(track, nums[i])
		(*used)[i] = true
		backtrack(nums, track, res, used)
		track = track[:len(track)-1]
		(*used)[i] = false
	}
}

func main() {
	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	fmt.Printf("res:%d", res)
}
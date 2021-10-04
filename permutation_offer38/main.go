package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {
	used, res, track := make([]bool, len(s)), []string{}, []rune{}
	nums := []rune(s)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	backtrack(nums, track, &res, &used)
	return res
}

func backtrack(nums []rune, track []rune, res *[]string, used *[]bool) {
	if len(track) == len(nums) {
		temp := make([]rune, len(track))
		copy(temp, track)
		*res = append(*res, string(temp))
		return
	}

	for i := 0; i < len(nums); i++ {
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
	s := "abc"
	fmt.Println(permutation(s))
}
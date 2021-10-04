package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sz := len(nums)
	index := make(map[int]int)
	for i, num := range nums {
		index[num] = i
	}
	for i := 0; i < sz; i++ {
		other := target - nums[i]
		// if _, ok := index[other] && index[other] != i {
		// 	return []int{i, index[other]}
		// }
		if _, ok := index[other]; ok && index[other] != i {
			return []int{i, index[other]}
		}
	}

	return []int{-1, -1}
}

func main() {
	nums := []int{2,7,11,15}
	res := twoSum(nums, 9)
	fmt.Println(res)
}
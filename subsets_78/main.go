package main

import "fmt"

func subsets(nums []int) [][]int {
	track, res := []int{}, [][]int{}
	backtrack(nums, 0, track, &res)
	return res
}

func backtrack(num []int, index int, track []int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)

	for i := index; i < len(num); i++ {
		track = append(track, num[i])
		backtrack(num, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	track, res := []int{}, [][]int{}
	backtrack(n, k, 1, track, &res)
	return res
}

func backtrack(target int, k int, index int, track []int, res *[][]int) {
	if target == 0 {
		if len(track) == k {
			temp := make([]int, len(track))
			copy(temp, track)
			*res = append(*res, temp)
		}
		return
	}

	for i := index; i < 10; i++ {
		track = append(track, i)
		backtrack(target-i, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
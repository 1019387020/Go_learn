package main

import "fmt"

func combine(n int, k int) [][]int {
	res, track := [][]int{}, []int{}
	backtrack(n, k, 1, track, &res)
	return res
}

func backtrack(n int, k int, index int, track []int, res *[][]int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := index; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func main() {
	n, k := 1, 1
	fmt.Println(combine(n, k))
}
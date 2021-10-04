package main

import (
	"fmt"
	"strconv"
	"time"
)

func getPermutation(n int, k int) string {
	used, track, res := make([]bool, n), []int{}, ""
	backtrack(n, &k, 0, &used, track, &res)
	return res
}

func backtrack(n int, k *int, index int, used *[]bool, track []int, res *string) {
	if n == index {
		*k--
		if *k == 0 {
			for _, v := range track {
				*res += strconv.Itoa(v+1)
			}
		}
		return
 	}

	for i := 0; i < n; i++ {
		if (*used)[i] {
			continue
		}
		
		track = append(track, i)
		(*used)[i] = true
		backtrack(n, k, index+1, used, track, res)
		track = track[:len(track)-1]
		(*used)[i] = false
	}
}

// 85.7607ms
// 57.6096ms
// 37.1093ms
func main() {
	strat := time.Now()
	fmt.Println(getPermutation(9, 78494))
	runTime := time.Since(strat)
	fmt.Println(runTime)
}

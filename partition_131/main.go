package main

import (
	"fmt"
)

func paritition(s string) [][]string {
	track, res := []string{}, [][]string{}
	backtrack(s, 0, track, &res)
	return res
}

func backtrack(s string, index int, track []string, res *[][]string) {
	if index == len(s) {
		temp := make([]string, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := index; i < len(s); i++ {
		if !isPalindrome(s, index, i) {
			continue
		}
		track = append(track, s[index : i+1])
		backtrack(s, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func isPalindrome(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "aba"
	fmt.Println(paritition(s))
}
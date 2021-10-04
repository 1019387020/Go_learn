package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	res := []string{}
	backtrack(s, wordDict, 0, "", &res)
	return res
}

func backtrack(s string, wordDict []string, index int, track string, res *[]string) {
	if index == len(s) {
		track = track[:len(track)-1]
		*res = append(*res, track)
		return
	}

	for i := 0; i < len(wordDict); i++ {
		if strings.HasPrefix(s[index:], wordDict[i]) {
			track = track + wordDict[i] + " "
			backtrack(s, wordDict, index+len(wordDict[i]), track, res)
			track = track[:len(track)-len(wordDict[i])-1]
		}
	}
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, wordDict))
}

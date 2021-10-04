package main

import "fmt"

// func ladderLength(beginWord string, endWord string, wordList []string) int {

// }

func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for i, word := range wordList {
		if word != beginWord {
			wordMap[word] = i
		}
	}
	return wordMap
}

func main() {
	wordList := []string{"hot","dot","dog","lot","log","cog"}
	beginWord := "hit"
	a := getWordMap(wordList, beginWord)
	fmt.Println(a)
}
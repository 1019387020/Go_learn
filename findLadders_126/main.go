package main

import (
	"fmt"
	"math"
	"time"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	used, track, res := make([]bool, len(wordList)), []string{}, [][]string{}
	have := false
	for _, v := range wordList {
		if v == endWord {
			have = true
			break
		}
	}
	if !have {
		return res
	}
	backtrack(beginWord, endWord, wordList, &used, track, &res)
	return finalRes(&res)
}

func backtrack(curWord string, endWord string, wordList []string, used *[]bool, track []string, res *[][]string) {
	if curWord == endWord {
		temp := make([]string, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(wordList); i++ {
		if !optional(curWord, wordList[i]) {
			continue
		}
		if (*used)[i] {
			continue
		}
		if len(track) == 0 {
			track = append(track, curWord)
		}

		track = append(track, wordList[i])
		(*used)[i] = true
		backtrack(wordList[i], endWord, wordList, used, track, res)
		track = track[:len(track)-1]
		(*used)[i] = false
	}
}

func optional(a string, b string) bool {
	num := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			num++
		}
	}
	return num == 1
}

func finalRes(res *[][]string) [][]string {
	minLen, finalRes := math.MaxInt64, [][]string{}
	for _, v := range *res {
		if len(v) < minLen {
			minLen = len(v)
		}
	}
	for _, v := range *res {
		if len(v) == minLen {
			finalRes = append(finalRes, v)
		}
	}
	return finalRes
}

func main() {
	start := time.Now()
	beginWord := "qa"
	endWord := "sq"
	wordList := []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
	fmt.Println(time.Since(start))
}

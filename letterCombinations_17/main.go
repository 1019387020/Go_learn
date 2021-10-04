package main

import "fmt"

var result []string
var dict = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	backtrack("", digits)
	return result
}

func backtrack(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}

	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		backtrack(res, digits)
		res = res[:len(res)-1]
	}
}

func main() {
	digits := "2"
	fmt.Println(letterCombinations(digits))
}

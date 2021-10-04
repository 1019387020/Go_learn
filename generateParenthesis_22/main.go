package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	res := []string{}
 	backtrack(n, n, "", &res)
	return res
}

func backtrack(left int, right int, track string, res *[]string) {
	/* 减枝 */
	if left > right {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	/* 减枝 */

	// 结束
	if left == 0 && right == 0 {
		*res = append(*res, track)
		return
	}

	/* 回溯模板 */ 
	track = track + "("
	backtrack(left-1, right, track, res)
	track = track[:len(track)-1]

	track = track + ")"
	backtrack(left, right-1, track, res)
	track = track[:len(track)-1]
	/* 回溯模板 */
}

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
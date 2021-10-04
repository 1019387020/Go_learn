package main

import "fmt"

// 优化版本，原地dp，time：O(n) space：O(1)
func longestValidParatheses(s string) int {
	left, right, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			res = max(res, 2*right)
		}
		if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			res = max(res, 2*left)
		}
		if left > right {
			left, right = 0, 0
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	//s := "()(()"
	s := "(()"
	//s := ")()())"
	//s := "()(())"
	res := longestValidParatheses(s)
	fmt.Println(res)
}

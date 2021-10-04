package main

import "fmt"

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	dp0, dp1, res := 1, 1, 0
	for i := 2; i < n+1; i++ {
		res = dp0 + dp1
		dp0 = dp1
		dp1 = res
	}
	return res
}

func main() {
	fmt.Println(climbStairs(7))
}
package main

import "fmt"

func trap(height []int) int {
	n, res := len(height), 0

	l_max, r_max := make([]int, n), make([]int, n)
	l_max[0], r_max[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		l_max[i] = max(height[i], l_max[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		r_max[i] = max(height[i], r_max[i+1])
	}
	for i := 1; i < n-1; i++ {
		res += min(l_max[i], r_max[i]) - height[i]
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(height))
}
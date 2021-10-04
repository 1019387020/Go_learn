package main

import "fmt"

func trap(height []int) int {
	n, res := len(height), 0
	
	for i := 1; i < n-1; i++ {
		left, right := 0, 0
		for j := i; j < n; j++ {
			right = max(right, height[j])
		}
		for j := i; j >= 0; j-- {
			left = max(left, height[j])
		}
		res += min(left, right) - height[i]
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
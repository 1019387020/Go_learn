package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	n, res := len(height), 0
	l_max, r_max := height[0], height[n-1]

	left, right := 0, n-1
	for left <= right {
		l_max = max(l_max, height[left])
		r_max = max(r_max, height[right])

		// res += min(l_max, r_max) - height[i]
		if l_max < r_max {
			res += l_max - height[left]
			left++
		} else {
			res += r_max - height[right]
			right--
		}
	}

	return res
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(height))
}
package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	n := len(nums)
	dpMax, dpMin := nums[0], nums[0]
	res := dpMax
	for i := 1; i < n; i++ {
		mx, mn := dpMax, dpMin
		dpMax = max(mx*nums[i], mn*nums[i], nums[i])
		dpMin = min(mx*nums[i], mn*nums[i], nums[i])
		if dpMax > res {
			res = dpMax
		}
	}
	return res
}

func max(x, y, z int) int {
	res := math.MinInt64
	if x > res {
		res = x
	}
	if y > res {
		res = y
	}
	if z > res {
		res = z
	}
	return res
}

func min(x, y, z int) int {
	res := math.MaxInt64
	if x < res {
		res = x
	}
	if y < res {
		res = y
	}
	if z < res {
		res = z
	}
	return res
}

func main() {
	nums := []int{2,-1,1,1}
	res := maxProduct(nums)
	fmt.Println(res)
}

package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	n := len(nums)
	dpMax, dpMin := make([]int, n), make([]int, n)
	dpMax[0], dpMin[0] = nums[0], nums[0]
	res := dpMax[0]
	for i := 1; i < n; i++ {
		dpMax[i] = max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
		dpMin[i] = min(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
		if dpMax[i] > res {
			res = dpMax[i]
		}
	}
	fmt.Println(dpMax)
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

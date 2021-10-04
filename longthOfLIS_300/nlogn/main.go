package main

import "fmt"

func longthOfLIS(nums []int) (int, []int) {
	top, piles := make([]int, len(nums)), 0
	for i := 0; i < len(nums); i++ {
		poker := nums[i]

		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] < poker {
				left = mid + 1
			} else if top[mid] > poker {
				right = mid
			} else {
				right = mid
			}
		}

		if left == piles {
			piles++
		}
		top[left] = poker
	}

	return piles, top[:piles]
}

func main() {
	nums := []int{0,1,0,3,2,3}
	fmt.Println(longthOfLIS(nums))
}
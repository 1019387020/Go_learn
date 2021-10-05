package main

import "fmt"

//时间复杂度：O(n)
//空间复杂度：O(1)
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return prices[0]
	}

	min, maxProfit := prices[0], 0
	for i := 1; i < n; i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return maxProfit
}

func main() {
	nums := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(nums))
}
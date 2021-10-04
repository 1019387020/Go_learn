package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitIntoFibonacci(num string) []int {
	if len(num) < 3 {
		return []int{}
	}

	res, isComplete := []int{}, false

	for firstEnd := 0; firstEnd < len(num)/2; firstEnd++ {
		if num[0] == '0' && firstEnd > 0 {
			break
		}
		first, _ := strconv.Atoi(num[:firstEnd+1])
		// 剪枝：结果数组中的每个数都要小于 2^31-1
		if first >= 1<<31 {
			break
		}
		for secondEnd := firstEnd+1; max(firstEnd, secondEnd-firstEnd) <= len(num)-secondEnd; secondEnd++ {
			if num[firstEnd+1] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			second, _ := strconv.Atoi(num[firstEnd+1 : secondEnd+1])
			// 剪枝：结果数组中的每个数都要小于 2^31-1
			if second >= 1<<31 {
				break
			}
			recursiveCheck(num, first, second, secondEnd+1, &res, &isComplete)
		}
	}

	return res
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func recursiveCheck(num string, x1 int, x2 int, left int, res *[]int, isComplete *bool) {
	if x1 >= 1<<31 || x2 >= 1<<31 {
		return 
	}
	
	if left == len(num) {
		if !*isComplete {
			*isComplete = true
			*res = append(*res, x1)
			*res = append(*res, x2) 
		}
		return
	}

	if strings.HasPrefix(num[left:], strconv.Itoa(x1+x2)) && !*isComplete {
		*res = append(*res, x1)
		recursiveCheck(num, x2, x1+x2, left+len(strconv.Itoa(x1+x2)), res, isComplete)
		return
	}
	
	// >= 0 切片索引异常
	if len(*res) > 0 && !*isComplete {
		*res = (*res)[:len(*res)-1]
	}
}

func main() {
	num := "539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"
	fmt.Println(splitIntoFibonacci(num))
}

package main

import "fmt"

func grayCode(n int) []int {
	var length uint = 1 << uint(n)
	res := make([]int, length)
	for i := uint(0); i < length; i++ {
		res[i] = int(i ^ (i >> 1))
	}
	return res
}

func main() {
	fmt.Printf("%d\n", grayCode(3))
}
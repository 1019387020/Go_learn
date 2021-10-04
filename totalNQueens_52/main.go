package main

import "fmt"

func totalNQueens(n int) int {
	rol, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, 0
	backtrack(n, 0, &rol, &dia1, &dia2, &row, &res)
	return res
}

func backtrack(n int, index int, rol *[]bool, dia1 *[]bool, dia2 *[]bool, row *[]int, res *int) {
	if index == n {
		*res++
		return
	}

	for i := 0; i < n; i++ {
		if (*rol)[i] || (*dia1)[index+i] || (*dia2)[index-i+n-1] {
			continue
		}

		*row = append(*row, i)
		(*rol)[i] = true
		(*dia1)[index+i] = true
		(*dia2)[index-i+n-1] = true
		backtrack(n, index+1, rol, dia1, dia2, row, res)
		*row = (*row)[:len(*row)-1]
		(*rol)[i] = false
		(*dia1)[index+i] = false
		(*dia2)[index-i+n-1] = false
	}
}

func main() {
	n := 1
	fmt.Println(totalNQueens(n))
}
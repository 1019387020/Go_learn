package main

import (
	"fmt"	
	"time"
)


func solveNQueens(n int) [][]string {
	col, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, [][]string{}
	backtrack(n, 0, &col, &dia1, &dia2, &row, &res)
	return res
}

func backtrack(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *[][]string) {
	if index == n {
		*res = append(*res, generateBoard(n, row))
		return
	}

	for i := 0; i < n; i++ {
		// 错误： if (*col)[i] && (*dia1)[index+i] && (*dia2)[index-i+n-1] {continue}
		                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          
		// if (*col)[i] {
		// 	continue
		// }
		// if (*dia1)[index+i] {
		// 	continue
		// }
		// if (*dia2)[index-i+n-1] {
		// 	continue
		// }
		if (*col)[i] || (*dia1)[index+i] || (*dia2)[index-i+n-1] {continue}

		// (*row)[index] = i
		*row = append(*row, i)
		(*col)[i] = true
		(*dia1)[index+i] = true
		(*dia2)[index-i+n-1] = true
		backtrack(n, index+1, col, dia1, dia2, row, res)
		//(*row)[index] = 0
		*row = (*row)[:len(*row)-1]
		(*col)[i] = false
		(*dia1)[index+i] = false
		(*dia2)[index-i+n-1] = false
	}
}

func generateBoard(n int, row *[]int) []string {
	board := []string{}
	res := ""
	
	for i := 0; i < n; i++ {
		res += "."
	}
	for i := 0; i < n; i++ {
		// board[i] = res
		board = append(board, res)
	}

	for i := 0; i < n; i++ {
		temp := []byte(board[i])
		temp[(*row)[i]] = 'Q'
		board[i] = string(temp) 
	}

	return board
}

func main() {
	start := time.Now()
	n := 4
	fmt.Println(solveNQueens(n))
	end := time.Since(start)
	fmt.Println(end)
}
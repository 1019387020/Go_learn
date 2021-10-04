package main

import "fmt"

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	// visited [][]bool ：记录字符是否被使用，避免重复选择
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	// 逐个遍历 board 中的字符
	for i, v := range board {
		for j := range v {
			// 判断能够找到 word 序列
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

// 判断 board 中 (x,y) 位置的字符是否与 word 中第 index 个字符相同
func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}

	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}

	return false
}

// 判断坐标 (x,y) 四周的点是否在 board 内
func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}



func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "SEDE"
	fmt.Println(exist(board, word))
}

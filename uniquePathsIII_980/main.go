package main

import "fmt"

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func uniquePathsIII(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	res, empty, startx, starty, endx, endy := 0, 0, 0, 0, 0, 0
	for i, v := range grid {
		for j, vv := range v {
			switch vv {
			case 0:
				empty++
			case 1:
				startx, starty = i, j
			case 2:
				endx, endy = i, j
			}
		}
	}
	findUniquePathIII(grid, visited, empty+1, startx, starty, endx, endy, &res)
	return res
}

func findUniquePathIII(board [][]int, visited [][]bool, empty, startx, starty, endx, endy int, res *int) {
	if startx == endx && starty == endy {
		if empty == 0 {
			*res++
		}
		return
	}

	if board[startx][starty] >= 0 {
		visited[startx][starty] = true
		empty--
		//track = append(track, startx)
		//track = append(track, starty)
		for i := 0; i < 4; i++ {
			nx := startx + dir[i][0]
			ny := starty + dir[i][1]
			// 错误：索引溢出，先后顺序导致的错误。
			if isInBoard(board, nx, ny) && !visited[nx][ny] {
				findUniquePathIII(board, visited, empty, nx, ny, endx, endy, res)
			}
		}
		visited[startx][starty] = false
		//empty++
		//track = track[:len(track)-2]
	}
}

func isInBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func main() {
	board := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}
	fmt.Println(uniquePathsIII(board))
}
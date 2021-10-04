package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	track, res := []int{}, [][]int{}
	backtrack(candidates, target, 0, track, &res)
	return res
}

// backtrack() 中借助 target 来控制 track 数组中的元素和，进行实现结束条件;
// backtrack() 中借助 index 实现决策树每层的选择列表，进行实现列举出所有组合序列。
func backtrack(candidates []int, target int, index int, track []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			temp := make([]int, len(track))
			copy(temp, track)
			*res = append(*res, temp)
		}
		return
	}

	for i := index; i < len(candidates); i++ {
		// 剪枝
		if candidates[i] > target {
			continue
		}

		track = append(track, candidates[i])
		backtrack(candidates, target-candidates[i], i, track, res)
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(nums, target))
}

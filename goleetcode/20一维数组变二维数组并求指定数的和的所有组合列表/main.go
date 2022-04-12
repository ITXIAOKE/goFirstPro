package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	if candidates == nil || len(candidates) == 0 {
		return res
	}
	cur := []int{}
	backtrack(candidates, target, 0, cur, &res)
	return res
}

func backtrack(candidates []int, target int, start int, cur []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		backtrack(candidates, target-candidates[i], i, cur, res)
		cur = cur[:len(cur)-1]
	}
}

func main() {
	//[2 3 6 7]  中等
	//求7的和  【2 2 3】 【7】
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

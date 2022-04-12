package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	if nums == nil || len(nums) == 0 {
		return res
	}
	cur := []int{}                     //最开始的时候什么都没有,后续去放二维数组里面的值【123】
	visited := make([]bool, len(nums)) //这个是标记元素是否被访问过
	backtrack(nums, cur, visited, &res)
	return res
}

func backtrack(nums []int, cur []int, visited []bool, res *[][]int) {
	if len(nums) == len(cur) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		//add option
		cur = append(cur, nums[i])

		//backtrack
		visited[i] = true
		backtrack(nums, cur, visited, res)
		visited[i] = false

		//remove option
		cur = cur[:len(cur)-1]

	}

}

func main() { //40 99 119
	//输入【1 2 3】
	//输出  【 【123】【132】【213】【231】【312】【321】】
	aa := []int{1, 2, 3}
	fmt.Println(permute(aa))
}
